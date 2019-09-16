#!/usr/bin/env bash

mongorestore "tests/business/automatic_tests" --host 127.0.0.1 --port 27017 --drop --db automatic_tests \
&& mongorestore "tests/business/automatic_tests_expected_results" --host 127.0.0.1 --port 27017 --drop --db automatic_tests_expected_results \
&& mongorestore "tests/business/validation_test_1" --host 127.0.0.1 --port 27017 --drop --db validation_test_1 \
&& mongorestore "tests/business/Scenario_Rules_Test_DB" --host 127.0.0.1 --port 27017 --drop --db Scenario_Rules_Test_DB \
&& mongorestore "tests/business/automatic_ndo_base_db_test" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test \
&& mongorestore "tests/business/automatic_ndo_base_db_test_1" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test_1 \
&& mongorestore "tests/business/automatic_ndo_base_db_test_2" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test_2 \
&& mongorestore "tests/business/automatic_ndo_base_db_test_3" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test_3 \
&& mongorestore "tests/business/automatic_ndo_base_db_test_4" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test_4 \
&& mongorestore "tests/business/automatic_ndo_base_db_test_5" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_base_db_test_5 \
&& mongorestore "tests/business/automatic_ndo_network_rules_test" --host 127.0.0.1 --port 27017 --drop --db automatic_ndo_network_rules_test \
&& mongorestore "tests/business/automatic_random_test" --host 127.0.0.1 --port 27017 --drop --db automatic_random_test \
&& mongorestore "tests/business/random_test_1" --host 127.0.0.1 --port 27017 --drop --db random_test_1 \
&& mongorestore "tests/business/random_test_2" --host 127.0.0.1 --port 27017 --drop --db random_test_2 \
&& mongorestore "tests/business/random_test_3" --host 127.0.0.1 --port 27017 --drop --db random_test_3
