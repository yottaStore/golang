# Yottablock specification

Yottablock is the disk storage format


Header:
- 4 bit version
- 12 bit size
- 16 bit flags

Total: 4 byte

Body:
- 4084 byte body

Footer:
- 64 bit hash


Total: 4096 byte
Usable: 4084 byte

# Type of blocks

- Single record smaller than usable space: single block
- Single record bigger than usable space: big block
- Multiple records (columnar storage): linked blocks


## Single block

Size count bytes after header
Footer contains hash of the first 4086 bytes
(header + body)

## Big block

If a record is bigger than a single block, flag
`F_BIGBLOCK` is set to 1, and size refer to 
the number of blocks that should be read, 
starting from this block. Blocks are contiguous.

If size is greater than 4096 blocks (~ 16,72 MiB),
then flag `F_JUMBOBLOCK` is set. 

TODO: Jumboblock

## Linked blocks

To store multiple records in one list of blocks, set
`F_LINKEDBLOCKS`, which means a header will be 
stored as first record in the block, with
data to navigate the list.

```
type LinkedHeader struct {
    NextLinkedBlock BlockPointer
    FirstAggregatedBlock BlockPointer
    AggregatedBlockCount uint64
}
```

Problem:
Block A contains pointer to block B. Block B gets 
compactified. How do we update the pointer?
Solution:
Compactify before other edits