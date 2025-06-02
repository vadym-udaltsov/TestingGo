.PHONY: all test report serve

RESULTS_DIR=allure-results
REPORT_DIR=allure-report

all: test report serve

test:
	@if not exist $(RESULTS_DIR) mkdir $(RESULTS_DIR)
	@cmd /C "set ALLURE_RESULTS_PATH=$(RESULTS_DIR) && go test ./tests"

report:
	@allure generate $(RESULTS_DIR) --clean -o $(REPORT_DIR)

serve:
	@allure open $(REPORT_DIR)
