'use strict'

module.exports = parse

/**
 * @param {Array[string]} argv
 * @param {Set[string]} keys
 * @return {Map[string]string}
 */

function parse(argv, keys) {
  let m = new Map()

  for (let i = 0; i < argv.length; i++) {
    if (!argv[i].startsWith('-')) continue

    if (!keys.has(argv[i])) {
      console.error('invalid argv: %s', argv[i])
      return process.exit(1)
    }

    if (argv[i + 1] && !argv[i + 1].startsWith('-')) {
      m.set(argv[i], argv[i + 1])
    } else {
      m.set(argv[i], '')
    }
  }

  return m
}
