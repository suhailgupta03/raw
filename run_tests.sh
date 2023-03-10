WITHCOVERAGE=$1
HOME=$HOME
TEST_DB_PATH="${HOME}/temp_db"
rm -rf "$TEST_DB_PATH"
cd src/tests
go test -v
if [ "$WITHCOVERAGE" == "-cover" ]; then
  rm -rf "$TEST_DB_PATH"
  go test -v -cover
else
    echo "Skipping the coverage"
fi