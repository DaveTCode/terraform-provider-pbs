## 0.1.0 (Unreleased)

FEATURES:

ENHANCEMENTS:

* resource/pbs_server: Added `unset_attributes`, a set of attribute names to explicitly reset. Scalar attributes are reset to their PBS default and map attributes have all entries removed. The value must be known at plan time, and each reset is applied once (tracked in private state) so Terraform state keeps reflecting the real PBS value. Because omitting an attribute now preserves its imported value, this provides an explicit way to remove a previously set attribute ([#101](https://github.com/DaveTCode/terraform-provider-pbs/issues/101)).

BUG FIXES:

* resource/pbs_server: Attributes omitted from configuration are no longer implicitly unset after import. Server attributes are now `Optional`+`Computed` with `UseStateForUnknown` plan modifiers, so unspecified attributes retain their imported values and a minimal import followed by `terraform plan` is clean and non-destructive ([#101](https://github.com/DaveTCode/terraform-provider-pbs/issues/101)).
* resource/pbs_server: Fixed `max_run_res` never being written to state on read, which caused a subsequent update to destructively unset every remote `max_run_res` entry ([#101](https://github.com/DaveTCode/terraform-provider-pbs/issues/101)).

