#ls -c | find . -type f -name "*.sh" | sed -e 's/\.[^.]*$//' | tr -d ./

find . -name '*.sh' | sed "s/.sh//g" | sed 's/\(.*\)\///g'