# go-notes

decode/marshall
---------------
Encoding is the process of transforming a piece of data into a “coded form”.
The “coded form” can be transformed back into the original message thanks to decoding.

Marshaling is the process of transforming an object into a storable representation using a specific format.
Unmarshaling is the reverse process.

An encoder operates on a stream of data, a marshaler operates on a data structure, an object.
Conclusion: we encode a stream of data, we marshal a data structure.
