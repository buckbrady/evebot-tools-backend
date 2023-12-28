# Changelog

## [1.2.1](https://github.com/buckbrady/evebot-tools-backend/compare/v1.2.0...v1.2.1) (2023-12-28)


### Bug Fixes

* remove bad column in market orders ([f7cffcd](https://github.com/buckbrady/evebot-tools-backend/commit/f7cffcddbc094961a155db7c12e886c4a22a9776))
* **tasks:** change market prices to standard queue ([f3dc7dd](https://github.com/buckbrady/evebot-tools-backend/commit/f3dc7dd8bca26e0f62ec617c2b72d788a1b51bdd))

## [1.2.0](https://github.com/buckbrady/evebot-tools-backend/compare/v1.1.0...v1.2.0) (2023-12-27)


### Features

* add market tasks ([ec1ef2a](https://github.com/buckbrady/evebot-tools-backend/commit/ec1ef2afaaa5b4f758e5957e8bf35035b2dbdff6))
* **tasks:** add realtime queue ([5d7bc26](https://github.com/buckbrady/evebot-tools-backend/commit/5d7bc2690f5f445359fc577617326a369848d82e))
* **tasks:** enable strict queue ([2b29c01](https://github.com/buckbrady/evebot-tools-backend/commit/2b29c0190d9db4bca9dd2cd9b0c78a7189e38432))

## [1.1.0](https://github.com/buckbrady/evebot-tools-backend/compare/v1.0.4...v1.1.0) (2023-12-25)


### Features

* **task:** add universe graphics sync ([a1f841c](https://github.com/buckbrady/evebot-tools-backend/commit/a1f841cb830cc5be102ca716db2290076f678374))
* **tasks:** add factions sync task ([6198ee6](https://github.com/buckbrady/evebot-tools-backend/commit/6198ee6d239e64a90a2a53b2831b699aa368918f))
* **tasks:** add system jumps task ([159e1ef](https://github.com/buckbrady/evebot-tools-backend/commit/159e1efb0f5beb34ec755d979d41b58ac82af78c))
* **tasks:** add system kills task ([de0caa8](https://github.com/buckbrady/evebot-tools-backend/commit/de0caa8a508cb5ac037dac533f8d21830c5c3758))
* **tasks:** add universe categories task ([cf14e77](https://github.com/buckbrady/evebot-tools-backend/commit/cf14e77196dc8404a5ba49ff39cde68a358ebbc1))
* **tasks:** add universe groups task ([cf14e77](https://github.com/buckbrady/evebot-tools-backend/commit/cf14e77196dc8404a5ba49ff39cde68a358ebbc1))


### Bug Fixes

* **scheduler:** add missing endpoints for new tasks ([5243097](https://github.com/buckbrady/evebot-tools-backend/commit/52430973f9d79edc7ea99f62290c0911f238c241))

## [1.0.4](https://github.com/buckbrady/evebot-tools-backend/compare/v1.0.3...v1.0.4) (2023-12-25)


### Bug Fixes

* **tasks:** fix static value in constellations ([b43a500](https://github.com/buckbrady/evebot-tools-backend/commit/b43a500d987acabbf8ca56f15c35d2d0bab78d04))

## [1.0.3](https://github.com/buckbrady/evebot-tools-backend/compare/v1.0.2...v1.0.3) (2023-12-25)


### Bug Fixes

* fix bad ctx call ([1b4da4d](https://github.com/buckbrady/evebot-tools-backend/commit/1b4da4d797eb0644facc8c657ca2c2dd6e438ca2))
* undo db query timeout ([#18](https://github.com/buckbrady/evebot-tools-backend/issues/18)) ([267a4f0](https://github.com/buckbrady/evebot-tools-backend/commit/267a4f0a975508f549fbf1bdeca2a606739c0dd3))

## [1.0.2](https://github.com/buckbrady/evebot-tools-backend/compare/v1.0.1...v1.0.2) (2023-12-25)


### Bug Fixes

* **database:** fix schema.hcl to use sql now() function as default ([#15](https://github.com/buckbrady/evebot-tools-backend/issues/15)) ([66aaeb3](https://github.com/buckbrady/evebot-tools-backend/commit/66aaeb35cf0a5def79bdef2e3b5a74bdac4e1c30))

## [1.0.1](https://github.com/buckbrady/evebot-tools-backend/compare/v1.0.0...v1.0.1) (2023-12-24)


### Bug Fixes

* add query timeout to db queries ([#12](https://github.com/buckbrady/evebot-tools-backend/issues/12)) ([dcfdf8b](https://github.com/buckbrady/evebot-tools-backend/commit/dcfdf8b0b17925b905214c08917b2a76bbb1500c))
* **tasks:** add check for bad typeID ([#14](https://github.com/buckbrady/evebot-tools-backend/issues/14)) ([d4affee](https://github.com/buckbrady/evebot-tools-backend/commit/d4affeeeef07a34bc4c8f4699a60e5c96833863b))

## 1.0.0 (2023-12-24)


### Features

* **scheduler:** add service scheduler ([41bf3c9](https://github.com/buckbrady/evebot-tools-backend/commit/41bf3c978ee5b1eb6f11af45132903d73ac71705))
* **worker:** add service scheduler ([41bf3c9](https://github.com/buckbrady/evebot-tools-backend/commit/41bf3c978ee5b1eb6f11af45132903d73ac71705))
* **zkillsync:** add sync service for zkillboard ([#7](https://github.com/buckbrady/evebot-tools-backend/issues/7)) ([41bf3c9](https://github.com/buckbrady/evebot-tools-backend/commit/41bf3c978ee5b1eb6f11af45132903d73ac71705))
