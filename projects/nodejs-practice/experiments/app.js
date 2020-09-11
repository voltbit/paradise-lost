const chalk = require('chalk')
const yargs = require('yargs')
const fs = require('fs')

function insertCassandraRow(argv) {
  console.log(chalk.blueBright(argv.columns))
  console.log(chalk.blueBright(argv.values))
}

yargs.command({
  command: 'add-row',
  describe: 'Add a new row in the database by declaring the columns and values',
  builder: {
    columns: {
      describe: 'column names',
      demandOption: true,
      type: 'array'
    },
    values: {
      describe: 'column values',
      demandOption: true,
      type: 'array'
    }
  }
})

function testJson() {
  const test = {
    key1: 'value1',
    key2: 'value2'
  }
  console.log(JSON.stringify(test))
  fs.writeFileSync('test.json', JSON.stringify(test))
}

function testArgs() {
  console.log(yargs.argv)
}

testJson()
