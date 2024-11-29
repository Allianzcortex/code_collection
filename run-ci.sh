cd advent_of_code 

# For solutions with go
dir_go_list="2021 2022 2023 2024"
for dir in $dir_go_list; do
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir" || exit 1

        if go test ; then
            echo "Tests passed in $dir"
        else
            echo "Tests failed in $dir"
        fi

        cd .. 
    fi
done

# run Python code
echo "==entering folder=="
cd "2024(python)"
echo "==running Python unitest=="
python -m unittest discover