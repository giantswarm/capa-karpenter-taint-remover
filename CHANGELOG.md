# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

## [1.0.1] - 2025-10-16

### Changed

- Use default catalog

## [1.0.0] - 2025-10-16

### Changed

- Not managing the CAPI taint anymore, it manages the `karpenter.sh/unregistered` karpenter taint that is applied by CAPI controllers when using a stale version for the patch request.

## [0.6.0] - 2025-05-26

### Changed

- Update Go version in Dockerfile to 1.23.

## [0.5.0] - 2025-05-22

### Changed

- Change node selector from `managed-by` to `karpenter.sh/registered`.

## [0.4.0] - 2024-10-10

### Changed

- Allow all taints in the daemonset

## [0.3.0] - 2024-03-19

### Changed

- Move to giantswarm catalog.

## [0.2.0] - 2024-03-19

### Added

- Wait to receive signal to quit gracefully.

## [0.1.0] - 2024-03-19

### Added

- First release.

[Unreleased]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.1...HEAD
[1.0.1]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.6.0...v1.0.0
[0.6.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.0.0...v0.1.0
