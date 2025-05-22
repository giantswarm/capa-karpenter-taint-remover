# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.0.0...v0.1.0
