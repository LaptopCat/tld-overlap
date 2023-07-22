from urllib.request import urlopen
from ast import literal_eval
from os import system

data = urlopen("https://data.iana.org/TLD/tlds-alpha-by-domain.txt").read().decode() # get all tlds for free
tlds = data.split("\n")[1:-1] # get all entries in a list and strip the shit
tlds = set(map(lambda a:a.lower().encode().decode("idna"), tlds)) # makes the tlds lowercase and decodes the punycode

with open("main.go") as file:
    data = file.read()
data = data.split("\n")
old_tlds = literal_eval("{" + "".join(data[8].split("{")[1:])) # .go ing insane
new = len(tlds) - len(old_tlds)
if new != 0:
    with open("main.go", "w") as file:
        data[8] = f"var tlds = [{len(tlds)}]string" + str(tlds).replace("'", '"')
        data = "\n".join(data)
        file.write(data)
    print(f"wrote {new} new tlds")
    print(", ".join(tlds - old_tlds))
    system("go build main.go")
else:
    print("no new tlds")
# patch old array of tlds and rebuild
