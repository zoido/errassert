# Changelog

## [1.0.0](https://github.com/zoido/errassert/compare/v0.5.2...v1.0.0) (2026-10-31)


### ⚠ BREAKING CHANGES

* github.com/zoido/errassert/grpcerrassert is now it's own module

### Features

* Extract grpcerrassert as a separated module ([cebedb2](https://github.com/zoido/errassert/commit/cebedb2594ccbf1054f755744c11eed9ffe3dc0f))

## [0.5.2](https://github.com/zoido/errassert/compare/v0.5.1...v0.5.2) (2026-10-28)


### Bug Fixes

* ⬆️ bump google.golang.org/grpc from 1.79.3 to 1.82.1 in the go_modules group across 1 directory ([#59](https://github.com/zoido/errassert/issues/59)) ([847ac9a](https://github.com/zoido/errassert/commit/847ac9a5bf7ebb359424f7e0280107db4609785d))
* ⬆️ bump google.golang.org/grpc from 1.82.1 to 1.83.1 in the go_modules group across 1 directory ([#60](https://github.com/zoido/errassert/issues/60)) ([31da658](https://github.com/zoido/errassert/commit/31da658d9a0e34c9d8373b25f027fac91d2d3a1d))
* **deps:** ⬆️ bump google.golang.org/grpc from 1.83.1 to 1.83.2 ([c5349f0](https://github.com/zoido/errassert/commit/c5349f09fa79397f4dfcfc2324f3317444bc6970))

## [0.5.1](https://github.com/zoido/errassert/compare/v0.5.0...v0.5.1) (2026-05-11)


### Bug Fixes

* ⬆️ bump google.golang.org/grpc from 1.56.3 to 1.79.3 in the go_modules group across 1 directory ([#51](https://github.com/zoido/errassert/issues/51)) ([37ade15](https://github.com/zoido/errassert/commit/37ade1594bc6794a4a23993845e27b2a4ff65db0))

## [0.5.0](https://github.com/zoido/errassert/compare/v0.4.0...v0.5.0) (2024-05-25)


### Features

* ✨ make `nil` `ErrorAssertion` a valid option ([#40](https://github.com/zoido/errassert/issues/40)) ([93decab](https://github.com/zoido/errassert/commit/93decabf59848194887e8e5405200de442ab23ed))

## [0.4.0](https://github.com/zoido/errassert/compare/v0.3.0...v0.4.0) (2024-04-12)


### Features

* ✨ support simple message equality check ([#35](https://github.com/zoido/errassert/issues/35)) ([bb2f95e](https://github.com/zoido/errassert/commit/bb2f95e8149bafb53e426571c9d61b251913c9d0))

## [0.3.0](https://github.com/zoido/errassert/compare/v0.2.0...v0.3.0) (2024-03-27)


### Features

* 📄 Open under MIT license ([4baa589](https://github.com/zoido/errassert/commit/4baa58934d97e677ce7322f245c28a5c9eca7bb1))

## [0.2.0](https://github.com/zoido/errassert/compare/v0.1.0...v0.2.0) (2024-03-15)


### Features

* ✨ allow to combine assertions with `Want` ([#23](https://github.com/zoido/errassert/issues/23)) ([c579a57](https://github.com/zoido/errassert/commit/c579a5717e7cfaebc734d14a78d48b87f54af2bd))
* ✨ support checking gRPC status errors ([#25](https://github.com/zoido/errassert/issues/25)) ([4ceacfc](https://github.com/zoido/errassert/commit/4ceacfc570d99addc645ba925539c93ce0439cda))

## [0.1.0](https://github.com/zoido/errassert/compare/v0.0.1...v0.1.0) (2024-03-13)


### Features

* 🚧 Basic nil/non-nil assertions ([#11](https://github.com/zoido/errassert/issues/11)) ([80a337c](https://github.com/zoido/errassert/commit/80a337cc54a6a19e1b1fcb583ab76142b6216239))
* 🚧 basic string based assertions ([#15](https://github.com/zoido/errassert/issues/15)) ([7678568](https://github.com/zoido/errassert/commit/767856893c1b3ba5454d454f2d2e113c016273bb))
* 🚧 Define the interface of the error assertion ([#8](https://github.com/zoido/errassert/issues/8)) ([f61f13c](https://github.com/zoido/errassert/commit/f61f13cf2d945e43c8c2b5f4cc790ec248e68e86))
* 🚧 go1.13 ErrorIs/ErrorAs assertions ([#12](https://github.com/zoido/errassert/issues/12)) ([155d565](https://github.com/zoido/errassert/commit/155d5651ffbb69b669e37f249184dfbf6f65a585))
