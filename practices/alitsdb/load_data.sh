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
        --urls          other venders' database urls
        --hosts         alitsdb hosts
        --port          alitsdb listenning port
        --do-load       whether to write the data to server actually"
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
    --hosts)
        _HOSTS=$2
        shift
        ;;
    --port)
        _PORT=$2
        shift
        ;;
    --do-load)
        _DOLOAD=$2
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
    _FORMAT="alitsdb"
fi
echo "Format is:$_FORMAT"

if [[ -z "$_BATCH_SIZE" ]]; then
    _BATCH_SIZE=1000
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

if [[ -z "$_HOSTS" ]]; then
    _HOSTS="127.0.0.1"
fi

if [[ -z "$_PORT" ]]; then
    _PORT=8242
fi

if [[ -z "$_DOLOAD" ]]; then
    _DOLOAD="true"
fi

rm -f ${_INPUT}/load_log
for file in ${_INPUT}/${_FORMAT}_seed_123_*
do
    echo ${file}
    if [ ${_FORMAT} = 'alitsdb' -o ${_FORMAT} = 'alitsdb-http' ]; then
        if [ ${_FORMAT} = 'alitsdb' ]; then
            cat ${file} | $GOPATH/bin/bulk_load_alitsdb -batch-size=${_BATCH_SIZE} -workers=${_WORKERS} -hosts=${_HOSTS} -port=${_PORT} -do-load=$_DOLOAD -json-format=false -viahttp=false >> ${_INPUT}/load_log 2>&1 &
        else
            cat ${file} | $GOPATH/bin/bulk_load_alitsdb -batch-size=${_BATCH_SIZE} -workers=${_WORKERS} -hosts=${_HOSTS} -port=${_PORT} -do-load=$_DOLOAD -json-format=true -viahttp=true >> ${_INPUT}/load_log 2>&1 &
        fi
    else
        cat ${file} | $GOPATH/bin/bulk_load_${_FORMAT} --batch-size=${_BATCH_SIZE} --workers=${_WORKERS} --urls=${_URLS} --do-db-create=false >> ${_INPUT}/load_log 2>&1 &
    fi
    sleep ${_SLEEP}
done
