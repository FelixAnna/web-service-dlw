app=$1
build="D:\code\github\keep-hands-on\important\dlw-app\build"
storage=dlwstorage916

echo "deploy $app frontend"
az storage blob upload-batch --account-name $storage -s $build -d '$web' --overwrite
