#!/usr/bin/env bash

while [ -n "$1" ]
do
  case "$1" in
    --help)

        echo "
        --help     Show this help
        --sec      how many sec's data do you want to generate
        --output   where do you want to generate
        --format   which gen cmd, such as es, influx"
        exit
        ;;
    --sec)
        _SEC=$2
        shift
        ;;
    --output)
        _OUTPUT=$2
        shift
        ;;
    --format)
        _FORMAT=$2
        shift
        ;;
    *)
        echo "$1 is not an option"
        exit
        ;;
  esac
  shift
done

if [[ -z "$_SEC" ]]; then
    _SEC=1
fi
echo "Seconds to generate：$_SEC "

if [[ -z "$_OUTPUT" ]]; then
    _OUTPUT="."
fi
mkdir -p ${_OUTPUT}
echo "The path of data：$_OUTPUT"

if [[ -z "$_FORMAT" ]]; then
    _FORMAT="es"
fi
echo "Format is：$_FORMAT"

start_time=1199116801
end_time=1199116802
echo $(date '+%Y-%m-%dT%H:%M:%SZ' -d @$start_time)
echo $(date '+%Y-%m-%dT%H:%M:%SZ' -d @$end_time)

for i in `seq 1 $_SEC`
do
        start_time=`expr $start_time + 1`
        end_time=`expr $end_time + 1`
        start_str=$(date '+%Y-%m-%dT%H:%M:%SZ' -d @$start_time)
        end_str=$(date '+%Y-%m-%dT%H:%M:%SZ' -d @$end_time)
        echo $start_str
        $GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=20000 --format=${_FORMAT}-bulk --timestamp-start=${start_str}  --timestamp-end=${end_str} | gzip > ${_OUTPUT}/es_seed_123_${start_time}.gz
done   
