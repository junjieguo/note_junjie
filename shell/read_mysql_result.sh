#!/bin/sh
QUERT_INFO=" test_filter"
QUERY_MYSQL="mysql"
QUERY_SQL="select ${QUERT_INFO} from  test_database.test_table limit 40;"
FILTER_INFO=$(${QUERY_MYSQL} -N -e "${QUERY_SQL}")
echo "$FILTER_INFO"


for filter in ${FILTER_INFO}
do


   
    echo -e "${filter}\n"


done
