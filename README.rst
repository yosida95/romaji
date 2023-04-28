romaji
======

romaji is a bi-directional translator for Katakana and Romaji.

Example
-------

.. code-block:: go

   package main

   import "github.com/yosida95/romaji"

   func main() {
   	kana := "トウキョウ"
   	roman := romaji.FromKanaHepburn(kana)
   	fmt.Printf("%q -> %q\n", kana, roman)

   	// Output:
   	// "トウキョウ" -> "TOKYO"
   }

Documentation
-------------

- `pkg.go.dev <https://pkg.go.dev/github.com/yosida95/romaji>`_


License
-------

`romaji`_ is distributed under the BSD 3-Clause license.
PLEASE READ ./LICENSE carefully and follow its clauses to use this software.

Author
------

yosida95_


.. _romaji: https://github.com/yosida95/romaji
.. _yosida95: https://yosida95.com/
