# Yottapack 

Yottapack is a bandwidth and cpu efficient 

- Size efficient
- Efficient streaming
- Query fields
- Schema or schemaless


# Format

## Header
- Intro byte
- 1+ size byte (needed for streaming)
- (optional): Heads encoding

### Heads encoding

Not needed if schemaless or streaming. Encode strings
representing the name of fields, (and nested fields?)

- 1 type byte (is it needed?)
- 1+ byte size
- N+ chars for value

## Data

- 1 type byte (optional)
- 1+ size byte
- N+ value byte

# Special bytes

## Intro

1 byte for versioning and flags

## Type byte

- 1 bit isArray
- 1 bit isNested
- 1 bit reserved
- 5 bit type (32 types)

### List of types

- Char
- Binary
- Bool
- Nil
- Integers ( 2 * 4 )
- BigInt
- BigDecimal
- Floats (2)
- Time ( 3 )
- Dynamic
- Lookup
- User defined ( 6 )
- Extended (2 bytes)

Total: 28

## Length encoding

1+ byte of u8. Like utf8

1st byte signal if you should read also next byte.

- 0 -> 128: 1 byte
- 128+1 -> 2^14-1: 2 bytes (first byte is 1, 8th byte is 0)
- 2^14+1 -> 2^21-1: 3 bytes (first byte is 1, 8th byte is 1, 
16th byte is 1)
- and so on....

## Separators

EOT or ETX used as separators. Escaped in strings