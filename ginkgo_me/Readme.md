The structure of the test is as following.
Please read comments in the following pseudo go code:

```
Describe("single struct - subject under test")
    // here we define subject under test
    // and all needed to initialize it mocks
    var subject impl
    var ... mocks...

    // In this BE (will run before each test) we initialize all the
    // mocks and the subject under test
    // This cause to all the tests to be purely separated one from another
    BeforeEach(...  subject = impl{...mocks...}

    // ensure all mocks are called after each test
    AfterEach(... mock... Finish())

    // Here we test specific subject's method
    // For each method we'll have separate D section
    Describe(".MethodName")

        // define inputs/output to the method under test
        // all imputs will be initialized in test cases (Context)
        // all outputs will be checked in asesrtion section (It)
        var in1, ... out1,...

        // Here we call the method under test (this will be executed ONCE
        // right before assertion section (It) - this means AFTER all
        // test cases (Context) sections, so all inputs will be surelty initialized
        JustBeforeEach( ... out1 = subject.MethodName(in1, in2))

        Context("when ...case description...")

            // Here we implementing test case description
            // This is done by:
            // - set input parameter to specific values
            // - set mock calls expectation
            // - sometimes (in gett) by storing data to redis/DB
            BeforeEach( ... case impl + mock.Expect()... }

            Context("when ...case description...")

                // same as note near previous BE
                BeforeEach({ ... case impl ... }

                // check method output only
                It(... {
                    Expect(output).To(HaveLen(4))
                    Expect(output[0]).To(Equal(...))
                    ...
                })
            })
        })
    })
})
```

