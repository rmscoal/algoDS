package com.example;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import org.junit.jupiter.api.Test;
import com.example.hackerrank.TryTesting;

class AppTest {
    @Test
    void testingPackageTest() {
        assertDoesNotThrow(() -> TryTesting.testingPackage());
    }
}

// Old Code
// import junit.framework.Test;
// import junit.framework.TestCase;
// import junit.framework.TestSuite;

// /**
// * Unit test for simple App.
// */
// public class AppTest
// extends TestCase {
// /**
// * Create the test case
// *
// * @param testName name of the test case
// */
// public AppTest(String testName) {
// super(testName);
// }

// /**
// * @return the suite of tests being tested
// */
// public static Test suite() {
// return new TestSuite(AppTest.class);
// }

// /**
// * Rigourous Test :-)
// */
// public void testApp() {
// assertTrue(true);
// }
// }