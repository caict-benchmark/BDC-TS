#!/usr/bin/env bash

while [ -n "$1" ]
do
  case "$1" in
    --help)

        echo "
        --help          Show this help
        --input         input data path
        --format        which gen cmd, such as es, influx
        --batch-size    batch size
        --workers       sync workers
        --sleep         sleep time
        --urls          database urls"
        exit
        ;;
    --input)
        _INPUT=$2
        shift
        ;;
    --format)
        _FORMAT=$2
        shift
        ;;
    --batch-size)
        _BATCH_SIZE=$2
        shift
        ;;
    --workers)
        _WORKERS=$2
        shift
        ;;
    --sleep)
        _SLEEP=$2
        shift
        ;;
    --urls)
        _URLS=$2
        shift
        ;;
    *)
        echo "$1 is not an option"
        exit
        ;;
  esac
  shift
done

if [[ -z "$_INPUT" ]]; then
    _INPUT="."
fi
echo "Generated data path:$_INPUT"

if [[ -z "$_FORMAT" ]]; then
    _FORMAT="es"
fi
echo "Format is:$_FORMAT"

if [[ -z "$_BATCH_SIZE" ]]; then
    _BATCH_SIZE=5000
fi

if [[ -z "$_WORKERS" ]]; then
    _WORKERS=5
fi

if [[ -z "$_SLEEP" ]]; then
    _SLEEP=0
fi

if [[ -z "$_URLS" ]]; then
    _URLS="http://localhost:9200"
fi

$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=1 --format=${_FORMAT}-bulk --timestamp-start=2017-01-01T00:00:00Z --timestamp-end=2017-01-01T00:00:01Z | $GOPATH/bin/bulk_load_${_FORMAT}  -workers 10
rm -f ${_INPUT}/load_log
for file in ${_INPUT}/${_FORMAT}_seed_123_*
do
    echo ${file}
    cat ${file} | $GOPATH/bin/bulk_load_${_FORMAT} --batch-size=${_BATCH_SIZE} --workers=${_WORKERS} --urls=${_URLS} --do-db-create=false >> ${_INPUT}/load_log 2>&1 &
    sleep ${_SLEEP}
done
