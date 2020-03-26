#!/usr/bin/env bash

while [ -n "$1" ]
do
  case "$1" in
    --help)

        echo "
        --help      Show this help
        --sec       how many sec's data do you want to generate
        --output    where do you want to generate(directory)
        --format    which gen cmd, such as es, influx
        --interval  interval for one file
        --scale     scale to generate
        --start-time epoch time(sec)"
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
    --interval)
        _INTERVAL=$2
        shift
        ;;
    --scale)
        _SCALE=$2
        shift
        ;;
    --start-time)
        _START_TIME=$2
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
echo "Seconds to generate:$_SEC "

if [[ -z "$_OUTPUT" ]]; then
    _OUTPUT="."
fi
mkdir -p ${_OUTPUT}
echo "The path of data:$_OUTPUT"

if [[ -z "$_FORMAT" ]]; then
    _FORMAT="alitsdb"
fi
echo "Format is:$_FORMAT"

if [[ -z "$_INTERVAL" ]]; then
    _INTERVAL=1
fi
echo "interval is:$_INTERVAL"

if [[ -z "$_SCALE" ]]; then
    _SCALE=20000
fi
echo "scale is:$_SCALE"

if (("$_SEC" < "$_INTERVAL")); then
    _INTERVAL=${_SEC}
fi

if [[ -z "$_START_TIME" ]]; then
    _START_TIME=1199116801
fi

start_time=${_START_TIME}
end_time=`expr ${start_time} + ${_SEC}`
echo "Generating data of the time range:"
echo $(date '+%Y-%m-%dT%H:%M:%SZ' -d @$start_time)
echo $(date '+%Y-%m-%dT%H:%M:%SZ' -d @$end_time)
file_num=`expr ${_SEC} / ${_INTERVAL}`

for i in `seq 1 ${file_num}`
do
        if [ $i -eq 1 ]
        then
            # start time as the ${_START_TIME}
            end_time=`expr ${start_time} + ${_INTERVAL}`
        else
            start_time=`expr ${start_time} + ${_INTERVAL}`
            end_time=`expr ${end_time} + ${_INTERVAL}`
        fi
        start_str=$(date '+%Y-%m-%dT%H:%M:%SZ' -d @$start_time)
        end_str=$(date '+%Y-%m-%dT%H:%M:%SZ' -d @$end_time)
        echo $start_str
        if [ ${_FORMAT} = 'alitsdb' -o ${_FORMAT} = 'alitsdb-http' ]; then
            echo "$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=${_SCALE} --format=${_FORMAT} --timestamp-start=${start_str}  --timestamp-end=${end_str} > ${_OUTPUT}/${_FORMAT}_seed_123_${start_time}"
            nohup $GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=${_SCALE} --format=${_FORMAT} --timestamp-start=${start_str}  --timestamp-end=${end_str} > ${_OUTPUT}/${_FORMAT}_seed_123_${start_time} 2>>/tmp/gen_logs &
        else
            echo "$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=${_SCALE} --format=${_FORMAT}-bulk --timestamp-start=${start_str}  --timestamp-end=${end_str} > ${_OUTPUT}/${_FORMAT}_seed_123_${start_time}.txt"
            nohup $GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=${_SCALE} --format=${_FORMAT}-bulk --timestamp-start=${start_str}  --timestamp-end=${end_str} > ${_OUTPUT}/${_FORMAT}_seed_123_${start_time}.txt 2>>/tmp/gen_logs &
        fi
done   
