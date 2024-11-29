cd advent_of_code 

# For solutions with go
for dir in *"(go)"; do
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir" || exit 1

        if go test ./...; then
            echo "Tests passed in $dir"
        else
            echo "Tests failed in $dir"
        fi
        
        # enter to hold folder
        cd .. 

for dir in *"(python)"; do
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir" || exit 1
    python -m unittest discover