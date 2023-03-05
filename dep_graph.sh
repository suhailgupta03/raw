# https://github.com/kisielk/godepgraph
# https://formulae.brew.sh/formula/graphviz
echo "Generating the dependency graph"
godepgraph src/main/main.go | dot -Tpng -o godepgraph.png
echo "Done 🚀"