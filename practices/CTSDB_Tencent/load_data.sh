#!/usr/bin/env bash

while [ -n "$1" ]
do
  case "$1" in
    --help)

        echo "
        --help     Show this help
        --input    input data path
        --format   which gen cmd, such as es, influx"
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
echo "Generated data path：$_INPUT"

if [[ -z "$_FORMAT" ]]; then
    _FORMAT="es"
fi
echo "Format is：$_FORMAT"

if [[ -z "$_BATCH_SIZE" ]]; then
    _BATCH_SIZE=5000
fi

if [[ -z "$_WORKERS" ]]; then
    _WORKERS=5
fi

if [[ -z "$_SLEEP" ]]; then
    _SLEEP=0
fi

for file in ${_INPUT}/es_seed_123_*
do
    echo ${file}
    cat ${_INPUT}/${file} | gunzip | $GOPATH/bin/bulk_load_${_FORMAT} --batch-size=${_BATCH_SIZE} --workers=${_WORKERS} > ${_INPUT}/load_log 2>&1 &
    sleep ${_SLEEP}
done
