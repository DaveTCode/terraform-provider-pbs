## 0.1.0 (Unreleased)

FEATURES:

BUG FIXES:

* resource/pbs_server: Attributes omitted from configuration are no longer implicitly unset after import. Server attributes are now `Optional`+`Computed` with `UseStateForUnknown` plan modifiers, so unspecified attributes retain their imported values and a minimal import followed by `terraform plan` is clean and non-destructive ([#101](https://github.com/DaveTCode/terraform-provider-pbs/issues/101)).
