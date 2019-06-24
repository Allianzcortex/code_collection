import nltk
from nltk.stem import PorterStemmer
from nltk import pos_tag
sentence = 'the little yellow dog barked at the cat'
grammar = ("'NP:{<DT>?<JJ>*<NN>}#NP'")
chunkParser = nltk.RegexpParser(grammar)
tagged = nltk.pos_tag(nltk.word_tokenize(sentence))
print(tagged)
tree = chunkParser.parse(tagged)
print(tree)