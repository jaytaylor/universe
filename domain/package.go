package domain

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/go/vcs"

	"jaytaylor.com/andromeda/pkg/unique"
)

// NewPackage turns a *vcs.RepoRoot into a new *Package.  If now is omitted or
// nil, the current time will be used.
func NewPackage(rr *vcs.RepoRoot, now ...*time.Time) *Package {
	pkg := &Package{
		Path:        rr.Root,
		Name:        "", // TODO: package name(s)???
		URL:         rr.Repo,
		VCS:         rr.VCS.Name,
		Data:        &PackageSnapshot{},
		ImportedBy:  map[string]*PackageReferences{},
		History:     []*PackageCrawl{},
		FirstSeenAt: orNow(now),
	}
	return pkg
}

func (pkg *Package) LatestCrawl() *PackageCrawl {
	l := len(pkg.History)
	if l == 0 {
		return nil
	}
	return pkg.History[l-1]
}

func (pkg *Package) MostlyEmpty() bool {
	return len(pkg.URL) == 0 && len(pkg.VCS) == 0 && pkg.Data == nil
}

// RepoRoot constructs a "fake" repository root, which may be adequate in some
// cases.
func (pkg Package) RepoRoot() *vcs.RepoRoot {
	rr := &vcs.RepoRoot{
		Repo: pkg.URL,
		Root: pkg.Path,
		VCS: &vcs.Cmd{
			Name: pkg.VCS,
		},
	}
	return rr
}

func (pkg Package) WebURL() string {
	expr := regexp.MustCompile(`^[^@]+@([^:]+):(.*)(?:\.git)?$`)
	if m := expr.FindAllStringSubmatch(pkg.URL, -1); len(m) > 0 {
		u := fmt.Sprintf("https://%v/%v", m[0][1], m[0][2])
		return u
	}
	return pkg.URL
}

// func NewSubPackage(path string, name string) {
// 	subPkg := &SubPackage{
// 		Path: path,
// 		Name: name,
// 	}
// 	return subPkg
// }

func (pkg *Package) Merge(other *Package) *Package {
	// TODO: When merging packages, it's going to be important to compare the the
	// imports this time to the previous imports, and generate some diffs to
	// update the newly referenced (or no longer referenced) packages.

	if pkg == nil {
		pkg = &Package{}
	}

	if pkg.ID == 0 && other.ID > 0 {
		pkg.ID = other.ID
	}
	if len(other.Name) > 0 {
		pkg.Name = other.Name
	}
	if len(other.Owner) > 0 {
		pkg.Owner = other.Owner
	}
	if len(other.Path) > 0 {
		pkg.Path = other.Path
	}
	if len(other.URL) > 0 {
		pkg.URL = other.URL
	}
	if len(other.VCS) > 0 {
		pkg.VCS = other.VCS
	}
	if other.Data != nil {
		if pkg.Data != nil {
			pkg.Data = pkg.Data.Merge(other.Data)
		} else {
			pkg.Data = other.Data
		}
	}
	if len(other.History) > 0 {
		// pkg.History = append(pkg.History, other.History...)
		pkg.History = other.History
	}
	/* DISABLED: THIS MAKES NO SENSE!!!
	if len(other.ImportedBy) > 0 {
		// pkg.ImportedBy = unique.StringsSorted(append(pkg.ImportedBy, other.ImportedBy...))
		for subPkgPath, wrapper := range other.ImportedBy {
			for _, ref := range wrapper.Refs {
				pkg.UpdateImportedBy(subPkgPath, ref)
			}
		}
	}*/
	return pkg
}

func (pkg *Package) MergePending(pendingRefs *PendingReferences) {
	if !strings.HasPrefix(pendingRefs.PackagePath, pkg.Path) {
		log.WithField("pkg", pkg.Path).WithField("pending-refs-pkg", pendingRefs.PackagePath).Warnf("Refusing to merge %v mismatched pending package references", len(pendingRefs.ImportedBy))
		return
	}

	for importPath, newRefs := range pendingRefs.ImportedBy {
		if pkg.ImportedBy == nil {
			pkg.ImportedBy = map[string]*PackageReferences{}
		}
		importPath = SubPackagePathNormalize(pkg.Path, importPath)
		if _, ok := pkg.ImportedBy[importPath]; !ok {
			pkg.ImportedBy[importPath] = NewPackageReferences()
		}
		for _, newRef := range newRefs.Refs {
			var found bool
			for _, ref := range pkg.ImportedBy[importPath].Refs {
				if ref.Path == newRef.Path {
					ref.Merge(newRef)
					found = true
					break
				}
			}
			if !found {
				pkg.ImportedBy[importPath].Refs = append(pkg.ImportedBy[importPath].Refs, newRef)
			}
		}
	}
}

func (pkg Package) Contains(path string) bool {
	if !strings.HasPrefix(path, pkg.Path) {
		return false
	}
	path = SubPackagePathNormalize(pkg.Path, path)
	for sub, _ := range pkg.Data.SubPackages {
		if sub == path {
			return true
		}
	}
	return false
}

func (ref *PackageReference) Merge(other *PackageReference) {
	if ref.Path != other.Path {
		log.WithField("ref", ref.Path).WithField("other", other.Path).Warn("Refusing to merge mismatched package reference")
		return
	}
	if !ref.Active && !other.Active {
		return
	}

	if other.LastSeenAt.After(*ref.LastSeenAt) {
		if ref.Active && !other.Active {
			ref.Active = false
		} else if !ref.Active && other.Active {
			ref.Active = true
		}
		ref.LastSeenAt = other.LastSeenAt
	}

	if other.FirstSeenAt.Before(*ref.FirstSeenAt) {
		ref.FirstSeenAt = other.FirstSeenAt
	}
}

/* This logic needs checking, or more likely to simply be tossed out completely.

// UpdateImportedBy behaves differently based on whether pkgRef.Active is true or
// false.
//
// When pkgRef.Active is true it adds or updates the timestamp on an importer.
//
// When pkgRef.Active is false it will be marked as inactive if the record is
// found, otherwise it'll be ignored.
func (pkg *Package) UpdateImportedBy(subPkgPath string, pkgRef *PackageReference) {
	subPkgPath = SubPackagePathNormalize(pkg.Path, subPkgPath)

	if wrapper, ok := pkg.ImportedBy[subPkgPath]; ok {
		var found bool
		for _, ref := range wrapper.Refs {
			if pkgRef.Path == ref.Path {
				if pkgRef.Active {
					ref.LastSeenAt = pkgRef.LastSeenAt
				} else {
					ref.Active = false
				}
				found = true
				break
			}
		}
		if !found {
			pkg.ImportedBy[subPkgPath].Refs = append(pkg.ImportedBy[subPkgPath].Refs, pkgRef)
		}
	} else {
		pkg.ImportedBy[subPkgPath] = &PackageReferences{Refs: []*PackageReference{pkgRef}}
	}
	// TODO: Sort alphabetically or by a sensible ranking metric.
}*/

// NormalizeSubPackageKeys runs the sub-package key normalizer on
// pkg.Data.SubPackages.
func (pkg *Package) NormalizeSubPackageKeys() {
	cleanedSubPkgs := map[string]*SubPackage{}
	for subPkgPath, subPkg := range pkg.Data.SubPackages {
		p := SubPackagePathNormalize(pkg.Path, subPkgPath)
		cleanedSubPkgs[p] = subPkg
	}
	pkg.Data.SubPackages = cleanedSubPkgs
}

type PackagePath struct {
	Name string
	Path string
}

// RepoName returns the tail end of the repository name.
// ".git" suffix will be removed.
func (pkg Package) RepoName() string {
	pieces := strings.Split(pkg.URL, "/")
	if len(pieces) == 0 {
		pieces = strings.Split(pkg.Path, "/")
	}
	name := pieces[len(pieces)-1]
	name = strings.TrimSuffix(name, ".git")
	return name
}

// ParentPaths returns one entry for each parent path of the package.
//
// Generate the path for a sub-package by passing in it's normalized or
// denormalized import path (it will be automatically normalized).
func (pkg Package) ParentPaths(subPkgPath ...string) []PackagePath {
	if len(subPkgPath) == 0 {
		subPkgPath = []string{""}
	} else if strings.HasPrefix(subPkgPath[0], pkg.Path) {
		subPkgPath[0] = SubPackagePathNormalize(pkg.Path, subPkgPath[0])
	}

	paths := []PackagePath{}
	pieces := strings.Split(pkg.Path, "/")

	if len(subPkgPath[0]) > 0 {
		pieces = append(pieces, strings.Split(subPkgPath[0], "/")...)
	}

	path := ""
	for i, piece := range pieces {
		if len(path) > 0 {
			path += "/"
		}
		path += piece
		if i < len(pieces)-1 {
			piece += "/"
		}
		pp := PackagePath{
			Name: piece,
			Path: path,
		}
		paths = append(paths, pp)
	}
	return paths
}

// SubPackagesPretty returns a map with all the sub-package keys denormalized.
func (pkg Package) SubPackagesPretty() map[string]*SubPackage {
	pretty := map[string]*SubPackage{}
	for k, v := range pkg.Data.SubPackages {
		// pretty[SubPackagePathDenormalize(pkg.Path, k)] = v
		pretty[k] = v
	}
	return pretty
}

func NewPackageReference(path string, now ...*time.Time) *PackageReference {
	pkgRef := &PackageReference{
		Path:        path,
		Active:      true,
		FirstSeenAt: orNow(now),
	}
	pkgRef.LastSeenAt = pkgRef.FirstSeenAt
	return pkgRef
}

func NewPackageReferences(refs ...*PackageReference) *PackageReferences {
	if refs == nil {
		refs = []*PackageReference{}
	}
	pkgRefs := &PackageReferences{
		Refs: refs,
	}
	return pkgRefs
}

func NewPackageCrawl(now ...*time.Time) *PackageCrawl {
	pc := &PackageCrawl{
		JobStartedAt: orNow(now),
		Data: &PackageSnapshot{
			SubPackages: map[string]*SubPackage{},
		},
	}
	pc.Data.CreatedAt = pc.JobStartedAt
	return pc
}

func (pc *PackageCrawl) AddMessage(msg string) {
	if pc.JobMessages == nil {
		pc.JobMessages = []string{}
	}
	// TODO: Perhaps change JobMessages to a map where the value is the number
	//       of occurrences.
	// Don't add duplicate messages.
	for _, jm := range pc.JobMessages {
		if jm == msg {
			return
		}
	}
	pc.JobMessages = append(pc.JobMessages, msg)
}

func (pc PackageCrawl) Duration() time.Duration {
	if pc.JobStartedAt == nil || pc.JobFinishedAt == nil {
		return time.Duration(0)
	}
	d := pc.JobFinishedAt.Sub(*pc.JobStartedAt)
	return d
}

func (snap *PackageSnapshot) AllImports() []string {
	imports := []string{}
	for _, subPkg := range snap.SubPackages {
		imports = append(imports, subPkg.Imports...)
	}
	imports = unique.StringsSorted(imports)
	return imports
}

// AllTestImports returns imports only found in package tests.
func (snap *PackageSnapshot) AllTestImports() []string {
	mainImports := map[string]struct{}{}
	for _, imp := range snap.AllImports() {
		mainImports[imp] = struct{}{}
	}
	testImports := []string{}
	for _, subPkg := range snap.SubPackages {
		for _, imp := range subPkg.TestImports {
			if _, ok := mainImports[imp]; !ok {
				testImports = append(testImports, imp)
			}
		}
	}
	testImports = unique.StringsSorted(testImports)
	return testImports
}

// CombinedImports returns all of both general and test imports.
func (snap *PackageSnapshot) CombinedImports() []string {
	imports := []string{}
	for _, subPkg := range snap.SubPackages {
		imports = append(imports, subPkg.Imports...)
		imports = append(imports, subPkg.TestImports...)
	}
	imports = unique.StringsSorted(imports)
	return imports
}

// Merge combines the information from a newer package snapshot into this one.
func (snap *PackageSnapshot) Merge(other *PackageSnapshot) *PackageSnapshot {
	// if snap == nil {
	// 	snap = &PackageSnapshot{}
	// }

	if snap.CreatedAt == nil && other.CreatedAt != nil {
		snap.CreatedAt = other.CreatedAt
	}
	if other.Repo != "" {
		snap.Repo = other.Repo
	}
	if other.SubPackages != nil {
		if snap.SubPackages == nil {
			snap.SubPackages = map[string]*SubPackage{}
		}
		// Search for sub-packages which no longer exist.
		for subPkgPath, subPkg := range snap.SubPackages {
			if otherSubPkg, ok := other.SubPackages[subPkgPath]; ok {
				// Accept updated fields.
				subPkg.LastSeenAt = otherSubPkg.LastSeenAt
				subPkg.Imports = otherSubPkg.Imports
				subPkg.TestImports = otherSubPkg.TestImports
				subPkg.Readme = otherSubPkg.Readme // TODO: Implement a de-duping method to avoid duplication.
			} else {
				// Sub-package is no longer available.
				subPkg.Active = false
			}
		}
		// Search and add new sub-packages.
		for otherSubPkgPath, otherSubPkg := range other.SubPackages {
			if _, ok := snap.SubPackages[otherSubPkgPath]; !ok {
				// Add new sub-package.
				snap.SubPackages[otherSubPkgPath] = otherSubPkg
			}
		}
	}
	if other.Commits != int32(0) {
		snap.Commits = other.Commits
	}
	if other.Branches != int32(0) {
		snap.Branches = other.Branches
	}
	if other.Tags != int32(0) {
		snap.Tags = other.Tags
	}
	if other.BytesVCS != uint64(0) {
		snap.BytesVCS = other.BytesVCS
	}
	if other.BytesTotal != uint64(0) {
		snap.BytesTotal = other.BytesTotal
	}
	if other.Stars != int32(0) {
		snap.Stars = other.Stars
	}
	if other.NumFiles != int32(0) {
		snap.NumFiles = other.NumFiles
	}
	if other.NumGoFiles != int32(0) {
		snap.NumGoFiles = other.NumGoFiles
	}
	if other.NumImports != int32(0) {
		snap.NumImports = other.NumImports
	}
	if other.NumTestImports != int32(0) {
		snap.NumTestImports = other.NumTestImports
	}

	snap.Sync()

	return snap
}

// Sync keeps calculated fields up to date so they don't become stale.
func (snap *PackageSnapshot) Sync() {
	snap.NumImports = int32(len(snap.AllImports()))
	snap.NumTestImports = int32(len(snap.AllTestImports()))
}

func (snap PackageSnapshot) Equals(other *PackageSnapshot) bool {
	if other == nil {
		return false
	}
	if snap.Branches != other.Branches {
		return false
	}
	if snap.BytesTotal != other.BytesTotal {
		return false
	}
	if snap.CommitHash != other.CommitHash {
		return false
	}
	if snap.Commits != other.Commits {
		return false
	}
	if snap.Repo != other.Repo {
		return false
	}
	if snap.Tags != other.Tags {
		return false
	}
	return true
}

func (snap PackageSnapshot) PrettyBytesTotal() string {
	pretty := humanize.Bytes(snap.BytesTotal)
	return pretty
}

func NewSubPackage(name string, now ...*time.Time) *SubPackage {
	subPkg := &SubPackage{
		Name:        name,
		Imports:     []string{},
		TestImports: []string{},
		FirstSeenAt: orNow(now),
	}
	subPkg.LastSeenAt = subPkg.FirstSeenAt
	return subPkg
}

func (subPkg *SubPackage) CombinedImports() []string {
	imports := make([]string, 0, len(subPkg.Imports)+len(subPkg.TestImports))
	for _, pkgPath := range subPkg.Imports {
		imports = append(imports, pkgPath)
	}
	for _, pkgPath := range subPkg.TestImports {
		imports = append(imports, pkgPath)
	}
	imports = unique.StringsSorted(imports)
	return imports
}

func (subPkg *SubPackage) MarkSeen(now ...*time.Time) {
	subPkg.LastSeenAt = orNow(now)
}

func (pkgRef *PackageReference) MarkSeen(now ...*time.Time) {
	pkgRef.LastSeenAt = orNow(now)
}

func SubPackagePathNormalize(pkgPath string, subPkgPath string) string {
	if subPkgPath == pkgPath {
		// NB: An empty string signifies the root package path.
		return ""
	}
	if !strings.HasPrefix(subPkgPath, pkgPath) {
		return subPkgPath
	}
	normalized := strings.Replace(subPkgPath, pkgPath+"/", "", 1)
	return normalized
}

func SubPackagePathDenormalize(pkgPath string, subPkgPath string) string {
	if subPkgPath == "" {
		return pkgPath
	}
	if strings.HasPrefix(subPkgPath, pkgPath) {
		return subPkgPath
	}
	denormalized := fmt.Sprintf("%v/%v", pkgPath, subPkgPath)
	return denormalized
}

func orNow(now []*time.Time) *time.Time {
	if len(now) == 0 {
		ts := time.Now()
		now = []*time.Time{&ts}
	}
	return now[0]
}
