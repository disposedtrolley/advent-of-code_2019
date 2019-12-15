(defproject aoc "0.1.0-SNAPSHOT"
  :description "Advent of Code 2019"
  :url ""
  :license {:name "MIT"}
  :dependencies [[org.clojure/clojure "1.10.0"]]
  :main ^:skip-aot aoc.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
