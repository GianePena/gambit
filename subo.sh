git add .
git commit -m "Ultimo Commit"
git push origin main

export GOOS=linux
export GOARCH=amd64

go build main.go
# Eliminar cualquier archivo ZIP existente
rm -f main.zip
# Crear un archivo ZIP de la compilación
zip main.zip main



