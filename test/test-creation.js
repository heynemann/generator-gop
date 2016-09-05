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
        services: ['redis'],
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

  //it('creates Makefile with right targets files', function (done) {
    //this.app.run({}, function () {
      //helpers.assertFileContent('Makefile', /list:/);
      //helpers.assertFileContent('Makefile', /no_targets__:/);
      //helpers.assertFileContent('Makefile', /setup:/);
      //helpers.assertFileContent('Makefile', /test:/);
      //helpers.assertFileContent('Makefile', /unit:/);
      //helpers.assertFileContent('Makefile', /redis:/);
      //helpers.assertFileContent('Makefile', /kill_redis:/);
      //helpers.assertFileContent('Makefile', /redis_test:/);
      //helpers.assertFileContent('Makefile', /kill_redis_test:/);
      //helpers.assertFileContent('Makefile', /mongo:/);
      //helpers.assertFileContent('Makefile', /kill_mongo:/);
      //helpers.assertFileContent('Makefile', /clear_mongo:/);
      //helpers.assertFileContent('Makefile', /mongo_test:/);
      //helpers.assertFileContent('Makefile', /kill_mongo_test:/);
      //helpers.assertFileContent('Makefile', /tox:/);
      //done();
    //});
  //});

  //it('creates setup.py with right elements', function (done) {
    //this.app.run({}, function () {
      //helpers.assertFileContent('setup.py', /from test_package import __version__/);
      //helpers.assertFileContent('setup.py', /name='test-package',/);
      //helpers.assertFileContent('setup.py', /    version=__version__,/);
      //helpers.assertFileContent('setup.py', /    description='an incredible python package',/);
      //helpers.assertFileContent('setup.py', /\nan incredible python package\n/);
      //helpers.assertFileContent('setup.py', /    keywords='test package',/);
      //helpers.assertFileContent('setup.py', /    author='Pablo Santiago Blum de Aguiar',/);
      //helpers.assertFileContent('setup.py', /    author_email='scorphus@gmail.com',/);
      //helpers.assertFileContent('setup.py', /    url='https:\/\/github.com\/someuser\/somepackage',/);
      //helpers.assertFileContent('setup.py', /    license='MIT',/);
      //helpers.assertFileContent('setup.py', /        'Programming Language :: Python :: 2.7',/);
      //helpers.assertFileContent('setup.py', /    include_package_data=False,/);
      //helpers.assertFileContent('setup.py', /        'pymongo',/);
      //helpers.assertFileContent('setup.py', /        'redis',/);
      //done();
    //});
  //});

});
