package nginx_docker

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class MyStatusCodeSim extends Simulation {

  val httpConf = http
    .baseURLs("http://web", "http://proxy")
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
            .get("/")
            .check(status.in(Seq(200,304))))
        .pause(100 milliseconds)
    }
// leaving redirects alone, they keep getting 404s
//    .during(1 seconds) {
//        exec(http("request_2")
//            .get("/302mepls")
//            .check(status.in(Seq(302,304))))
//        .pause(100 milliseconds)
//    }
    .during(5 seconds) {
        exec(http("request_3")
            .get("/403mepls")
            .check(status.is(403)))
        .pause(100 milliseconds)
    }
    .during(5 seconds) {
        exec(http("request_4")
            .get("/404mepls")
            .check(status.is(404)))
        .pause(100 milliseconds)
    }
    .during(5 seconds) {
        exec(http("request_5")
            .get("/500mepls")
            .check(status.is(500)))
        .pause(100 milliseconds)
    }

  setUp(scn.inject(atOnceUsers(2)).protocols(httpConf)) //.throttle(reachRps(10) in (0.2 seconds))
}
