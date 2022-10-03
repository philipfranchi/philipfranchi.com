yarn build
mkdir tmp
cp -r build tmp/public
rm -f go.zip
cp application.go go.mod tmp/
cd tmp
zip -r go.zip application.go public/*
cd ..
mv tmp/go.zip .
rm -rf tmp