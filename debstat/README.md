Solving time - 2h 15min excluding the README

---

### General thoughts about the code

The script downloads the required Contents file and stores it in a temporary file on the disk using functioanlity form urllib. The file is then read, table rows are filtered and converted into a hashmap which is later sorted and displayed.

- Chose to read the file line by line because it will scale with the file size and it is easier to do more advanced parsing on the content since the guildelines allow for free text at teh beginning of the file and to not enforce a particular table header.
- The code was not written with regard to python version compatibility, but it was a deliberate choice since it was not a requirement and I like to use python3 features (like f-strings). I am aware that the fstrings will not work in a python2 interpreter and probably also the urllib imports.
- Chose to use a return code '1' for http errors in case the tool is used as part of other scripts or along other system tools that can make use of it
- I assumed the encoding of files and locations are all UTF-8 encoded
- I was not sure how to handle the fact that the file can contain free text in the first lines, a simple assumption I've made is that no package will have file in the root directory and tehrefore one table entry must always contain at least one '/' character, a harder constrint would be for the package name to also contain at lease one '/' followed by alphanumerical characters
---

### Nitpicks

- Chose to not use any external libraries for parsing the file contents because the tool is more of a utility script and a data analysis program, otherwise I would have chosen to work with `pandas`.
- The URLs class is not an enum because I don't like that enums require to acces values via `.value` field so `URLs.file.value` instead of `URLs.file`
- Structured everything as a static class although it is not necessary because the tool could be used as a utility class/library

Other ideas (I did not get to implement or thought are unnecessary)

- use a cache folder - if the usecase allows for it, then it is better to cache the downloads for
a period of time in a local hidden folder in case the engineer/systrm runs the commadn a lot in a
short period of time
- add a `--verbose` param to show more info like length of downloaded file
- should look more into possible encodings of the data
- try to find a better parse rule for the header of the file and make sure that no packages can have files inside the root (so as to not exclude them form the list)
