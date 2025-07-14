import json

with open("block1.ipynb", encoding="utf-8") as f:
    nb = json.load(f)

# Удаляем метаданные jupytext
if "jupytext" in nb["metadata"]:
    del nb["metadata"]["jupytext"]

with open("block1.ipynb", "w", encoding="utf-8") as f:
    json.dump(nb, f, indent=1, ensure_ascii=False)
