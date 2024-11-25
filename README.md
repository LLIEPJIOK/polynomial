## Polynomial project

This program implements various operations with polynomials over the finite field \( F_{2^n} \). It supports:

- **Addition**: Fast XOR operation between the coefficients of two polynomials.
- **Multiplication**: Polynomial multiplication with reduction modulo a given irreducible polynomial.
- **Division**: Computes the quotient and remainder of polynomial division.
- **Exponentiation**: Raises a polynomial to a specified power with modular reduction.
- **Inverse Polynomial**: Finds the multiplicative inverse using the extended Euclidean algorithm.
- **Irreducibility Check**: Verifies if a polynomial is irreducible over \( F_{2^n} \).

This program is useful for applications in cryptography, error-correcting codes, and finite field algebra.

---

## Run

1.  Clone the repository:

    ```bash
    git clone git@github.com:LLIEPJIOK/polynomial.git
    ```

2.  Navigate to the project folder:

		```bash
		cd polynomial
		```

3.  Run the program:

    ```bash
    go run cmd/polynomial/main.go
    ```
