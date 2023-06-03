# [gomock](https://github.com/golang/mock)
- Common Configuration

| option  | description |
|---------|-------------|
| `-source` | A file containing interfaces to be mocked.|
| `-destination`| A file to which to write the resulting source code. If you don't set this, the code is printed to standard output.|
| `-package`| The package to use for the resulting mock class source code. If you don't set this, the package name is `mock_` concatenated with the package of the input file.|