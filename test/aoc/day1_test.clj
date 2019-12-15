(ns aoc.day1-test
  (:require [clojure.test :refer :all]
            [aoc.day1 :refer :all]))

(deftest test-fuel
  (def tests [{:input 12 :expected 2}
              {:input 14 :expected 2}
              {:input 1969 :expected 654}
              {:input 100756 :expected 33583}])

  (doseq [t tests]
    (is (= (fuel (get t :input)) (get t :expected))))
)
