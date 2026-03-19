# Changelog

## 1.5.0 (2026-03-19)

Full Changelog: [v1.4.2...v1.5.0](https://github.com/with-ours/platform-cli/compare/v1.4.2...v1.5.0)

### Features

* **api:** api update ([08577d6](https://github.com/with-ours/platform-cli/commit/08577d6fc551dc14a7ea1129f8906cca403bd29a))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([25295d5](https://github.com/with-ours/platform-cli/commit/25295d55458cf3f1606a743d8b34460a8b965c5c))
* better support passing client args in any position ([cd7c596](https://github.com/with-ours/platform-cli/commit/cd7c596b12c3649143931d35ca30614706b7e5eb))
* improve linking behavior when developing on a branch not in the Go SDK ([da9eb75](https://github.com/with-ours/platform-cli/commit/da9eb75eae8991db1e58640b6b2b8e5350231d5e))
* improved workflow for developing on branches ([0337780](https://github.com/with-ours/platform-cli/commit/033778056a78504c4f1d500273e6e03b23e5b233))
* no longer require an API key when building on production repos ([b999f2e](https://github.com/with-ours/platform-cli/commit/b999f2ec5bbbbbd170d523d622853d26d509d2ba))


### Chores

* **internal:** tweak CI branches ([f50e126](https://github.com/with-ours/platform-cli/commit/f50e126bc3314adf7fc2dcf6e132aa2259b786a1))

## 1.4.2 (2026-03-14)

Full Changelog: [v1.4.1...v1.4.2](https://github.com/with-ours/platform-cli/compare/v1.4.1...v1.4.2)

### Bug Fixes

* only set client options when the corresponding CLI flag or env var is explicitly set ([6d6b690](https://github.com/with-ours/platform-cli/commit/6d6b690520220496e3bf6236cb2672ebbf480a28))

## 1.4.1 (2026-03-12)

Full Changelog: [v1.4.0...v1.4.1](https://github.com/with-ours/platform-cli/compare/v1.4.0...v1.4.1)

### Bug Fixes

* fix for test cases with newlines in YAML and better error reporting ([0357fb2](https://github.com/with-ours/platform-cli/commit/0357fb2d42c7c6b237a6c219e4763875e3cd301d))

## 1.4.0 (2026-03-09)

Full Changelog: [v1.3.0...v1.4.0](https://github.com/with-ours/platform-cli/compare/v1.3.0...v1.4.0)

### Features

* **api:** api update ([7e15e55](https://github.com/with-ours/platform-cli/commit/7e15e555aaa25994bdee5ad19d7958f5a9e84b5c))

## 1.3.0 (2026-03-07)

Full Changelog: [v1.2.0...v1.3.0](https://github.com/with-ours/platform-cli/compare/v1.2.0...v1.3.0)

### Features

* **api:** manual updates ([31f0338](https://github.com/with-ours/platform-cli/commit/31f03383ca97470f74b60c1d78a84fd829a7bf82))


### Bug Fixes

* fix for encoding arrays with `any` type items ([546318a](https://github.com/with-ours/platform-cli/commit/546318a28946678300413e4583a5843562c0750a))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([799faea](https://github.com/with-ours/platform-cli/commit/799faeadf92f7f54b187eb54da9b5770c1cc9bef))

## 1.2.0 (2026-03-05)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/with-ours/platform-cli/compare/v1.1.0...v1.2.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([22dfe08](https://github.com/with-ours/platform-cli/commit/22dfe08431cca5cef1120c519d04d2d3c97c9b39))
* add support for file downloads from binary response endpoints ([7a981e5](https://github.com/with-ours/platform-cli/commit/7a981e5ba9df6e9ad738be2db20ffdb30e432047))
* **api:** api update ([07ec71b](https://github.com/with-ours/platform-cli/commit/07ec71bff3aa3c867ab29bac24b35aeb2dc0d964))
* **api:** api update ([af9d689](https://github.com/with-ours/platform-cli/commit/af9d689a29c51b81be093e3c23dc27131710f0b6))
* **api:** api update ([16d9929](https://github.com/with-ours/platform-cli/commit/16d9929409f1d77750fdaaca6e6cd4dc97137408))
* support passing required body params through pipes ([4126af6](https://github.com/with-ours/platform-cli/commit/4126af6a2ccffe6dff89bd976cf113f3fb1aa5fd))


### Bug Fixes

* add missing client parameter flags to test cases ([83072ed](https://github.com/with-ours/platform-cli/commit/83072ed08d34ae08e609ed4df0bd73ef9bcbb920))
* avoid printing usage errors twice ([6fd4cbf](https://github.com/with-ours/platform-cli/commit/6fd4cbf4bb58ec2daa8e81bcca5a20e8f04b38c2))
* more gracefully handle empty stdin input ([504e889](https://github.com/with-ours/platform-cli/commit/504e88903e014548694a2a7190e7c7166f33b730))
* pin formatting for headers to always use repeat/dot formats ([5ae0d55](https://github.com/with-ours/platform-cli/commit/5ae0d55e53ac9b236601f8756b6eac02900c436e))


### Chores

* **internal:** codegen related update ([aeb2d05](https://github.com/with-ours/platform-cli/commit/aeb2d05c9da3bf4cb5a451aa3e83dcb8ba000270))
* **internal:** codegen related update ([557d33c](https://github.com/with-ours/platform-cli/commit/557d33c0d71dc4bb74fa98cd396f55e7a498fd97))
* **internal:** codegen related update ([b442265](https://github.com/with-ours/platform-cli/commit/b4422655bdd7f314bb76d82c327dd66d2fd99623))
* **test:** do not count install time for mock server timeout ([8b9a668](https://github.com/with-ours/platform-cli/commit/8b9a668a7bcccb79944b1e9bde2eb1b01daaec5a))
* update readme with better instructions for installing with homebrew ([0951857](https://github.com/with-ours/platform-cli/commit/09518570a3c01da7e3ece1d2788e73923a0824c0))
* zip READMEs as part of build artifact ([d4ca0e7](https://github.com/with-ours/platform-cli/commit/d4ca0e7cadabdcb00a9a43e1c44c61de4a4766ca))

## 1.1.0 (2026-02-20)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/with-ours/platform-cli/compare/v1.0.0...v1.1.0)

### Features

* **api:** manual updates ([c173240](https://github.com/with-ours/platform-cli/commit/c1732404df3586a850c380597bc3e559ce304183))

## 1.0.0 (2026-02-20)

Full Changelog: [v0.0.1...v1.0.0](https://github.com/with-ours/platform-cli/compare/v0.0.1...v1.0.0)

### Features

* **api:** manual updates ([43c8fc1](https://github.com/with-ours/platform-cli/commit/43c8fc1d4d3889be2dd5070b21229740e82c4780))


### Chores

* configure new SDK language ([4c3284c](https://github.com/with-ours/platform-cli/commit/4c3284cd8c9c23d9651c6018d41026da4713a259))
