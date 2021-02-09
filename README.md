# dsp0134

This package is an example and should not be used as it is unsupported and will remain so.  There may be a supported version in the future, but it may not be at this location.

Example UUID package for the DSP0134 encoding of UUIDs

This package is a wrapper around the [github.com/google/uuid](https://github.com/google/uuid) package that supports the non-standard UUID encoding as defined by the [DMTF DSP0134](https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.4.0.pdf) specification.

This specifcation encodes `"00112233-4455-6677-8899-AABBCCDDEEFF"` as
```
33 22 11 00 55 44 77 66 88 99 AA BB CC DD EE FF
```
rather than the standard encoding of
```
00 11 22 33 44 55 66 77 88 99 AA BB CC DD EE FF
```

The `uuid.UUID` method in this package can be used to call inspection methods from the standard uuid package.  For example:
```
v := u.UUID().Version
```
Will return the proper version of u.
