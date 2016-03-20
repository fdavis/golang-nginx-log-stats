package computerdatabase

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class MySimulation extends Simulation {

  val httpConf = http
    .baseURLs("http://web", "http://proxy")
//    .baseURL("http://web")
    .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8") // Here are the common headers
    .doNotTrackHeader("1")
    .acceptLanguageHeader("en-US,en;q=0.5")
    .acceptEncodingHeader("gzip, deflate")
    .userAgentHeader("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0")

  val headers_10 = Map("Content-Type" -> "application/x-www-form-urlencoded") // Note the headers specific to a given request

  // A scenario to help verify golang stats
  val scn = scenario("Test Status Codes")
    .during(5 seconds) {
        exec(http("request_1")
            .get("/"))
    }
    .during(5 seconds) {
        exec(http("request_2")
            .get("/403mepls"))
    }
    .during(5 seconds) {
        exec(http("request_3")
            .get("/404mepls"))
    }
    .during(5 seconds) {
        exec(http("request_4")
            .get("/500mepls"))
    }

  setUp(scn.inject(atOnceUsers(5)).protocols(httpConf)).throttle(reachRps(10) in (0.2 seconds))
}
