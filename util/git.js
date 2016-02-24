
'use strict'

const Path = require('path')
const resolve = Path.resolve
const fs = require('fs')
const join = Path.join

module.exports = {
  dirRepos,
  isRepo,
  notDir,
}

/**
 * @param {string} path
 * @param {number} depth
 * @return {array[string]}
 */
function dirRepos(path, depth) {
  if (!--depth) return []
  if (notDir(path)) return []
  if (isRepo(path)) return [path]

  const names = fs.readdirSync(path)
  let repos = []

  names.forEach(n => {
    let p = join(path, n)

    if (notDir(p)) return
    if (isRepo(p)) return repos.push(p)

    repos = repos.concat(dirRepos(p, depth))
  })

  return repos
}

/**
 * @param {string} dir
 * @return {boolean}
 */
function isRepo(dir) {
  let p = join(dir, '.git')

  try {
    fs.statSync(p)
  } catch (e) {
    return false
  }

  return true
}

/**
 * @param {string} path
 * @return {boolean}
 */
function notDir(path) {
  let isDir

  try {
    isDir = fs.statSync(path).isDirectory()
  } catch (e) {
    return true
  }

  return !isDir
}
