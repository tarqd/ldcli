# Changelog

## [0.12.1](https://github.com/launchdarkly/ldcli/compare/v0.12.0...v0.12.1) (2024-05-10)


### Miscellaneous Chores

* release 0.12.1 ([#265](https://github.com/launchdarkly/ldcli/issues/265)) ([33f606d](https://github.com/launchdarkly/ldcli/commit/33f606d394fd67b5a353d669ead3dfad98de68fc))

## [0.12.0](https://github.com/launchdarkly/ldcli/compare/v0.11.0...v0.12.0) (2024-05-10)


### Features

* break out required and optional flags in subcommand help ([#262](https://github.com/launchdarkly/ldcli/issues/262)) ([f3075bf](https://github.com/launchdarkly/ldcli/commit/f3075bf2d5144b67bbafd2afeb1fbad69f193124))
* hide resources ([#254](https://github.com/launchdarkly/ldcli/issues/254)) ([af22b5d](https://github.com/launchdarkly/ldcli/commit/af22b5d501c08d48bb466a99b39c868c6f2a2679))
* show pagination ([#261](https://github.com/launchdarkly/ldcli/issues/261)) ([9a71cfe](https://github.com/launchdarkly/ldcli/commit/9a71cfe6fededaa2fe2e89970fe69ba91a86c055))


### Bug Fixes

* fix go sdk instructions ([#253](https://github.com/launchdarkly/ldcli/issues/253)) ([2fe1c8a](https://github.com/launchdarkly/ldcli/commit/2fe1c8a4bd283fe3bd703ad80fc4fff36687b5f4))
* Fix plaintext output for a resource with only and ID ([#263](https://github.com/launchdarkly/ldcli/issues/263)) ([e9139c0](https://github.com/launchdarkly/ldcli/commit/e9139c0df2a9e166a0ec2977b2559627e645775f))
* remove http client timeout for resources ([#259](https://github.com/launchdarkly/ldcli/issues/259)) ([56cbcdb](https://github.com/launchdarkly/ldcli/commit/56cbcdb54daa1e2823db64c7c03418fdc9a36994))
* typo in SDK setup steps ([#257](https://github.com/launchdarkly/ldcli/issues/257)) ([99379aa](https://github.com/launchdarkly/ldcli/commit/99379aae617f66d1678b4f465b8088aaec0fffd1))

## [0.11.0](https://github.com/launchdarkly/ldcli/compare/v0.10.0...v0.11.0) (2024-05-07)


### Features

* create and generate template from template data for teams ([#238](https://github.com/launchdarkly/ldcli/issues/238)) ([bf2f0a1](https://github.com/launchdarkly/ldcli/commit/bf2f0a10aeba241ef32bf7c37b8a58ce1f2567d4))
* generate remaining resources commands from openapi spec ([#244](https://github.com/launchdarkly/ldcli/issues/244)) ([e78e32b](https://github.com/launchdarkly/ldcli/commit/e78e32b5f53b13ac17a2575c479f00cdebc590b4))
* throttle flag toggle ([#243](https://github.com/launchdarkly/ldcli/issues/243)) ([3b88329](https://github.com/launchdarkly/ldcli/commit/3b88329d3e2dc4a57f8b70cef9078b40a1103b00))
* track help cmd ([#245](https://github.com/launchdarkly/ldcli/issues/245)) ([1ebc398](https://github.com/launchdarkly/ldcli/commit/1ebc398a184239a62df7b1ac5e33f123f342c6c1))


### Bug Fixes

* remove old members cmd ([#249](https://github.com/launchdarkly/ldcli/issues/249)) ([27c72ce](https://github.com/launchdarkly/ldcli/commit/27c72ce35c4369ea1a19166b9ed96a99f8543913))

## [0.10.0](https://github.com/launchdarkly/ldcli/compare/v0.9.0...v0.10.0) (2024-05-02)


### Features

* generate teams operation data from openapi spec ([#226](https://github.com/launchdarkly/ldcli/issues/226)) ([e96fb54](https://github.com/launchdarkly/ldcli/commit/e96fb54fef91a08df3e6b3d3cf690fcf15c0dd94))
* generic api request function ([#218](https://github.com/launchdarkly/ldcli/issues/218)) ([0141d07](https://github.com/launchdarkly/ldcli/commit/0141d07c02bfdf60b8ee5bad2f5981348180242d))
* update sdk instructions ([#230](https://github.com/launchdarkly/ldcli/issues/230)) ([2909424](https://github.com/launchdarkly/ldcli/commit/29094241577240f7899391807cd8ab5bd2e531b2))

## [0.9.0](https://github.com/launchdarkly/ldcli/compare/v0.8.1...v0.9.0) (2024-04-30)


### Features

* Add additional help text with missing access-token ([#219](https://github.com/launchdarkly/ldcli/issues/219)) ([b74053c](https://github.com/launchdarkly/ldcli/commit/b74053cbda60f8450b6c943d7f55bc3cc8eb649e))
* add hardcoded operation command with no body ([#211](https://github.com/launchdarkly/ldcli/issues/211)) ([c27e904](https://github.com/launchdarkly/ldcli/commit/c27e90431aaecd5e2c76fd47009c5b59e2d246d0))
* Add valid config fields to its help ([#217](https://github.com/launchdarkly/ldcli/issues/217)) ([ffa9fb3](https://github.com/launchdarkly/ldcli/commit/ffa9fb32274ddc54d9086d0c79a4f518660a83a6))
* added cmd completed to commands ([#200](https://github.com/launchdarkly/ldcli/issues/200)) ([bf0f6aa](https://github.com/launchdarkly/ldcli/commit/bf0f6aa51ee92acae053a6ee53b5fb49693dc782))
* allow users to opt out of analytics tracking ([#206](https://github.com/launchdarkly/ldcli/issues/206)) ([e782a43](https://github.com/launchdarkly/ldcli/commit/e782a431eb7cc73da863c8736c2fed6a82c26d7c))
* Create an --output/-o flag for JSON or plain text responses ([#195](https://github.com/launchdarkly/ldcli/issues/195)) ([96474cd](https://github.com/launchdarkly/ldcli/commit/96474cdaaba4175d1b0d9c20c4d1ece8b43ef7ee))
* hardcoded resource cmds ([#203](https://github.com/launchdarkly/ldcli/issues/203)) ([b8dc52a](https://github.com/launchdarkly/ldcli/commit/b8dc52a0fbde1c6f7c67ea102ca2893031ac0de3))
* output flag all commands ([#201](https://github.com/launchdarkly/ldcli/issues/201)) ([1670cae](https://github.com/launchdarkly/ldcli/commit/1670cae8239555920587af7fb619e564b93af0fe))
* output in config ([#209](https://github.com/launchdarkly/ldcli/issues/209)) ([e246cbc](https://github.com/launchdarkly/ldcli/commit/e246cbc692634334bee7d77403daca142f4afc8f))
* plaintext success response ([#210](https://github.com/launchdarkly/ldcli/issues/210)) ([82244ed](https://github.com/launchdarkly/ldcli/commit/82244edd2983da668480b9f6f0f6784c25e48938))
* show successful resource delete message ([#212](https://github.com/launchdarkly/ldcli/issues/212)) ([c1c3c1a](https://github.com/launchdarkly/ldcli/commit/c1c3c1a2f74216db7f1c2e8cdc78123d5c9ffc49))
* track cli command run events ([#189](https://github.com/launchdarkly/ldcli/issues/189)) ([fd98b42](https://github.com/launchdarkly/ldcli/commit/fd98b421cf09c7ee50266d9d73a654101ad11d7d))
* track cli setup step started event ([#215](https://github.com/launchdarkly/ldcli/issues/215)) ([25b9f2e](https://github.com/launchdarkly/ldcli/commit/25b9f2e6c7577ff300cc09f493ed16dfb4be6917))
* track sdk selected on setup ([#221](https://github.com/launchdarkly/ldcli/issues/221)) ([2e3445c](https://github.com/launchdarkly/ldcli/commit/2e3445c7487babb65b8b50a27cd05f866510de76))
* track setup flag toggle event ([#222](https://github.com/launchdarkly/ldcli/issues/222)) ([3b408cc](https://github.com/launchdarkly/ldcli/commit/3b408cca643f5a2a70dc320305a30b5d13a0df77))


### Bug Fixes

* config action output ([#225](https://github.com/launchdarkly/ldcli/issues/225)) ([889d8a2](https://github.com/launchdarkly/ldcli/commit/889d8a268c8c3fa6ec20857f4f7745c447deb7eb))
* Fix data flag JSON error handling ([#214](https://github.com/launchdarkly/ldcli/issues/214)) ([a469c0c](https://github.com/launchdarkly/ldcli/commit/a469c0c8b7d7854feda27c40a3bfcacf1b1b2986))
* remove get projects limit ([68d0bf0](https://github.com/launchdarkly/ldcli/commit/68d0bf04607133293d87108cf67e9354c0394c96))

## [0.8.1](https://github.com/launchdarkly/ldcli/compare/v0.8.0...v0.8.1) (2024-04-22)


### Bug Fixes

* Don't find/replace FLAG_KEY in SDK instructions ([#193](https://github.com/launchdarkly/ldcli/issues/193)) ([39ea1ce](https://github.com/launchdarkly/ldcli/commit/39ea1ce37e25619ad7059a6dccd54837fdb6c7b3))

## [0.8.0](https://github.com/launchdarkly/ldcli/compare/v0.7.0...v0.8.0) (2024-04-18)


### Features

* add get subcommand to flag ([#180](https://github.com/launchdarkly/ldcli/issues/180)) ([19443ab](https://github.com/launchdarkly/ldcli/commit/19443ab96420cd93af277ece0a9f069d11bbd375))
* display current flag state on toggle flag page ([#183](https://github.com/launchdarkly/ldcli/issues/183)) ([cfb3c1f](https://github.com/launchdarkly/ldcli/commit/cfb3c1fb206e7838ef4237167e242e8785a8a08e))


### Bug Fixes

* Change config file location ([#185](https://github.com/launchdarkly/ldcli/issues/185)) ([cc987c4](https://github.com/launchdarkly/ldcli/commit/cc987c46c4550a0fba2800470fd0961901ba1a61))
* remove gif and reference from readme ([#184](https://github.com/launchdarkly/ldcli/issues/184)) ([db6f378](https://github.com/launchdarkly/ldcli/commit/db6f378fc7bf1db6a1b7dc5979ec9064ac4030b8))
* Show help when running `ldcli config` ([#186](https://github.com/launchdarkly/ldcli/issues/186)) ([92a3e51](https://github.com/launchdarkly/ldcli/commit/92a3e51d96d25e02162f79235f0a82a0db67da98))

## [0.7.0](https://github.com/launchdarkly/ldcli/compare/v0.6.0...v0.7.0) (2024-04-16)


### Features

* Add config command ([#175](https://github.com/launchdarkly/ldcli/issues/175)) ([c1feb53](https://github.com/launchdarkly/ldcli/commit/c1feb53e3499af5b8c9d57cd49f76d97205b4429))
* add remaining SDK instructions ([#164](https://github.com/launchdarkly/ldcli/issues/164)) ([284669d](https://github.com/launchdarkly/ldcli/commit/284669d181e825fa5d0918cdac13ca0bb12ca9ff))
* change flag toggle success message for mobile sdks ([#156](https://github.com/launchdarkly/ldcli/issues/156)) ([37c2d6d](https://github.com/launchdarkly/ldcli/commit/37c2d6d9c96fb207d56deca4d7906fe5c89b5f6a))
* publish to npm ([#159](https://github.com/launchdarkly/ldcli/issues/159)) ([461467f](https://github.com/launchdarkly/ldcli/commit/461467f1b0b94037d15e2553a9a08983272ca9ea))
* support env vars ([#171](https://github.com/launchdarkly/ldcli/issues/171)) ([b0380ca](https://github.com/launchdarkly/ldcli/commit/b0380caafa35579a55420ef29f8e31fb7fcd9105))
* Use new sdk instructions instead of hello READMEs ([#152](https://github.com/launchdarkly/ldcli/issues/152)) ([6826a5c](https://github.com/launchdarkly/ldcli/commit/6826a5c9b61705f1e8f340bbd22779330cd8ee34))


### Bug Fixes

* embed instructions files & show error during show sdk step ([#174](https://github.com/launchdarkly/ldcli/issues/174)) ([ae07b46](https://github.com/launchdarkly/ldcli/commit/ae07b461f93e1f3c797138d0b89cdb7a6d16f297))
* remove mouse wheel support from show sdk & show scroll options in help view ([#161](https://github.com/launchdarkly/ldcli/issues/161)) ([99015b1](https://github.com/launchdarkly/ldcli/commit/99015b113077a881eb5886e86b5de8b9b82af4d2))
* remove side borders from show SDK viewport ([#162](https://github.com/launchdarkly/ldcli/issues/162)) ([d111c61](https://github.com/launchdarkly/ldcli/commit/d111c61c1f6792e2d1e88f45646a0828ec72dca6))
* space above pagination dots and pagination bug ([#155](https://github.com/launchdarkly/ldcli/issues/155)) ([adbc53c](https://github.com/launchdarkly/ldcli/commit/adbc53c0ee39e0f9de6c54e7e8e4111f8665e6aa))

## [0.6.0](https://github.com/launchdarkly/ldcli/compare/v0.5.0...v0.6.0) (2024-04-09)


### Features

* make sdk instructions scrollable ([#150](https://github.com/launchdarkly/ldcli/issues/150)) ([8055927](https://github.com/launchdarkly/ldcli/commit/805592776b336b3eb8bf0e99247025ee170d1b30))
* show success message after creating a flag ([#134](https://github.com/launchdarkly/ldcli/issues/134)) ([f817856](https://github.com/launchdarkly/ldcli/commit/f817856158d8205fbd11952ac2a664137900566c))

## [0.5.0](https://github.com/launchdarkly/ldcli/compare/v0.4.0...v0.5.0) (2024-04-05)


### Features

* show help text ([#138](https://github.com/launchdarkly/ldcli/issues/138)) ([69d7f5e](https://github.com/launchdarkly/ldcli/commit/69d7f5ee03579fc6c3cb34fb00fb6f07ba90bf84))

## [0.4.0](https://github.com/launchdarkly/ldcli/compare/v0.3.0...v0.4.0) (2024-04-05)


### Features

* publish docker image ([#99](https://github.com/launchdarkly/ldcli/issues/99)) ([a294ce0](https://github.com/launchdarkly/ldcli/commit/a294ce063e608a6dbee5ff8a2fb32f329d643083))


### Bug Fixes

* create flag prompt should be on the same line as text ([#135](https://github.com/launchdarkly/ldcli/issues/135)) ([52aa92b](https://github.com/launchdarkly/ldcli/commit/52aa92b647710f86f531b69afb3be50baf452a8d))

## [0.3.0](https://github.com/launchdarkly/ldcli/compare/v0.2.0...v0.3.0) (2024-04-04)


### Features

* Can go back to choose SDK page from show SDK instructions ([#120](https://github.com/launchdarkly/ldcli/issues/120)) ([6900bc6](https://github.com/launchdarkly/ldcli/commit/6900bc6301bf668830ca029867f6d9de4c555d85))


### Bug Fixes

* add instructions to continue from show sdk step ([#117](https://github.com/launchdarkly/ldcli/issues/117)) ([4baf2d0](https://github.com/launchdarkly/ldcli/commit/4baf2d0b2850b3979267f1d2ab964b5a295a27e3))
* fix the step count ([#121](https://github.com/launchdarkly/ldcli/issues/121)) ([1943c45](https://github.com/launchdarkly/ldcli/commit/1943c4590587f0a3b84e60c6a1628a2ed37254b6))
* remove id from goreleaser homebrew build config ([#125](https://github.com/launchdarkly/ldcli/issues/125)) ([56538cd](https://github.com/launchdarkly/ldcli/commit/56538cd34bb3ea908e4622a110ab6f486c288b86))

## [0.2.0](https://github.com/launchdarkly/ldcli/compare/v0.1.0...v0.2.0) (2024-04-03)


### Features

* add ldcli formula to homebrew-tap ([#108](https://github.com/launchdarkly/ldcli/issues/108)) ([1d638dc](https://github.com/launchdarkly/ldcli/commit/1d638dc0bf23f1e9fab96718a0a0058a74284dc0))
* Add more help in error message ([#72](https://github.com/launchdarkly/ldcli/issues/72)) ([6221983](https://github.com/launchdarkly/ldcli/commit/62219838f1c815695f6529252ded80170f267dcf))
* add sdk instructions step to quickstart ([#91](https://github.com/launchdarkly/ldcli/issues/91)) ([bf4aba6](https://github.com/launchdarkly/ldcli/commit/bf4aba61651c7bc359157e39abceb9b4bf7b101e))
* add toggle flag step ([#111](https://github.com/launchdarkly/ldcli/issues/111)) ([9cd4018](https://github.com/launchdarkly/ldcli/commit/9cd401883f0b2a9b1d84fde095658316091ecd3d))
* add toggle on and off aliases to update flag ([#82](https://github.com/launchdarkly/ldcli/issues/82)) ([7b6c6f1](https://github.com/launchdarkly/ldcli/commit/7b6c6f1e160598353ab0ed0e93cbc1bd3dd04360))
* alias command to invite members ([#84](https://github.com/launchdarkly/ldcli/issues/84)) ([7002866](https://github.com/launchdarkly/ldcli/commit/7002866263f510637546c3fea02608083cfb689b))
* better flag create error handling ([#85](https://github.com/launchdarkly/ldcli/issues/85)) ([708925b](https://github.com/launchdarkly/ldcli/commit/708925b4978f85539fffca513c84434f00b5d08f))
* Create choose SDK view ([#89](https://github.com/launchdarkly/ldcli/issues/89)) ([7518423](https://github.com/launchdarkly/ldcli/commit/7518423da3712fa2986fe20935428d7336d21b81))
* create command to get an environment ([#96](https://github.com/launchdarkly/ldcli/issues/96)) ([19b4ede](https://github.com/launchdarkly/ldcli/commit/19b4ede40d836855896c06105e7f0fb01d7e6161))
* create command to invite members ([#68](https://github.com/launchdarkly/ldcli/issues/68)) ([e1a2ca5](https://github.com/launchdarkly/ldcli/commit/e1a2ca5d09415143cc0673674d83d29fc1a99e17))
* create prod-ready quickstart command ([#75](https://github.com/launchdarkly/ldcli/issues/75)) ([7768bfa](https://github.com/launchdarkly/ldcli/commit/7768bfae10842292e615c5da8b18e3ee94066049))
* pass in optional role flag when using members invite ([#90](https://github.com/launchdarkly/ldcli/issues/90)) ([b51470d](https://github.com/launchdarkly/ldcli/commit/b51470d54d6c63fcc478b64770a2c6e91a2539fb))
* setting version dynamically ([#83](https://github.com/launchdarkly/ldcli/issues/83)) ([7e4e794](https://github.com/launchdarkly/ldcli/commit/7e4e7943a425689e342e34c55a613bf02b7c44c2))
* Show error along with message and quit if applicable ([#87](https://github.com/launchdarkly/ldcli/issues/87)) ([f076a59](https://github.com/launchdarkly/ldcli/commit/f076a59772b85b095411f63f35e1c66bee0304d0))


### Bug Fixes

* fix err messages for toggle flag command ([#112](https://github.com/launchdarkly/ldcli/issues/112)) ([9b385e9](https://github.com/launchdarkly/ldcli/commit/9b385e9a0816ba1677aaf170b1186f86f69d786a))
* members create cmd should invite multiple members ([#103](https://github.com/launchdarkly/ldcli/issues/103)) ([c575248](https://github.com/launchdarkly/ldcli/commit/c575248c6ab058622dd6e55f2b312656743e4cff))
* rebind projKey flag to update subcommand ([#62](https://github.com/launchdarkly/ldcli/issues/62)) ([1f0f898](https://github.com/launchdarkly/ldcli/commit/1f0f898511716becdf57ecd90a81ee84c56b1217))
* remove create flag placeholder text ([#114](https://github.com/launchdarkly/ldcli/issues/114)) ([3e8624d](https://github.com/launchdarkly/ldcli/commit/3e8624d6e70c45285143a5bbded5c917fff829bc))

## 0.1.0 (2024-03-22)

### Miscellaneous Chores

* release initial version
