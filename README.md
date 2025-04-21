# Usage

In your terminal, use go run ./cmd/shellspy

Type "exit" without the quotation marks to exit

# Goals

- [x] Goal 1: creating a command from a string
- [x] Goal 2: looping and reading input
- [x] Goal 3: executing a command and getting its output
    Supports:
    - ls
    - echo
    - mv
- [x] Goal 4: basic shell
- [x] Goal 5: transcript recording

# Stretch Goals

- [ ] Implement a remote shell
- [ ] Add auth
- [ ] Support pipes: |
- [ ] Support input redirection operator <

# Acknowledgements

Based on Bitfield Consulting's [shellspy project](https://github.com/bitfield/shellspy)

# Feedback

1. Place tests in shellspy_test package rather than directly in the shellspy package so that the tests have to interact with the public shellspy API. This gives a better sense of how actual users will interact with the API