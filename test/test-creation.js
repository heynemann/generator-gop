/*global describe, beforeEach, it */
'use strict';
const path = require('path')
const helpers = require('yeoman-test')
const assert = require('yeoman-assert');
const chai = require('chai')
const expect = chai.expect;
const gopath = process.env.GOPATH
const exec = require('child_process').exec;

const runBinary = function(command, cb) {
  exec(command, cb)
}

const getCoverage = function(stdout) {
  const re = /(?:coverage[:]\s*(\d+[.]\d+)[%])/gi
  let coverage = 0
  let coverageCount = 1
  let result = re.exec(stdout)
  while (result) {
    coverage += parseFloat(result[1])
    coverageCount++
    result = re.exec(stdout)
  }

  return coverage / Math.max(coverageCount - 1, 1)
}

describe('gop generator', function () {
  describe('main command', function() {
    describe('without services', function() {
      beforeEach(function () {
        const self = this
        expect(gopath).to.not.be.null
        return helpers.
          run(path.join(__dirname, '../app')).
          on('ready', function (generator) {
            self.generator = generator
          }).
          inDir(path.join(gopath, 'src/github.com/heynemann/test-package')).
          withPrompts({
            packageName: 'test-package',
            description: 'an incredible python package',
            keywords: 'test package',
            authorName: 'Bernardo Heynemann',
            authorEmail: 'heynemann@gmail.com',
            url: 'https://github.com/heynemann/test-package',
            license: 'MIT',
            services: [],
          }).toPromise()
      })

      it('creates expected files', function () {
        var expected = [];
        for (var i=0; i < this.generator.packageContents.length; i++) {
          const item = this.generator.packageContents[i]
          expected.push(item.target)
        }

        assert.file(expected)
      })

      it('runs setup', function(done) {
        this.timeout(15000);
        runBinary('make setup', function(error, stdout, stderr) {
          expect(error).to.be.null
          done()
        });
      })

      it('runs tests', function(done) {
        this.timeout(30000);
        runBinary('make setup test', function(error, stdout, stderr) {
          expect(error).to.be.null

          expect(getCoverage(stdout)).to.equal(100.0)
          done()
        });
      })

      it('runs cross-compiling', function(done) {
        this.timeout(15000);
        runBinary('make cross', function(error, stdout, stderr) {
          expect(error).to.be.null

          assert.file([
            'bin/test-package-linux-i386',
            'bin/test-package-linux-x86_64',
            'bin/test-package-darwin-i386',
            'bin/test-package-darwin-x86_64',
          ])

          done()
        })
      })
    })
  })
})
