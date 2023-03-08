# https://github.com/kisielk/godepgraph
# https://formulae.brew.sh/formula/graphviz
# If you want to ignore standard library packages entirely, use the -s flag:
echo "Generating the dependency graph"
godepgraph -s src/main/main.go | dot -Tpng -o godepgraph.png
echo "Done 🚀"