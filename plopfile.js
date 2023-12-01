require('dotenv').config();

const { Client } = require('pg');
const { promisify } = require('util');
const fs = require('fs');
const writeFileAsync = promisify(fs.writeFile);

const client = new Client({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_DATABASE,
  password: process.env.DB_PASSWORD,
  port: process.env.DB_PORT, 
});

const fetchTablesAndData = async () => {
  try {
    await client.connect();
    const tableResult = await client.query(
      "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'"
    );

    const tables = tableResult.rows.map(row => row.table_name);
    const allTableData = {};

    for (const table of tables) {
      const result = await client.query(`SELECT * FROM ${table}`);
      allTableData[table] = result.rows;
    }

    return allTableData;
  } catch (err) {
    console.error('Error fetching data from database:', err);
    throw err;
  } finally {
    await client.end();
  }
};

const fetchDataFromDB = async (tableName) => {
  try {
    await client.connect();
    const result = await client.query(`SELECT * FROM ${tableName}`);
    return result.rows; // Assuming your_table exists and fetches data
  } catch (err) {
    console.error('Error fetching data from database:', err);
    throw err;
  } finally {
    await client.end();
  }
};

module.exports = function (plop) {
  plop.setGenerator('fetchFromDB', {
    description: 'Generate files based on data from database by using input',
    prompts: [
      {
        type: 'input',
        name: 'tableName',
        message: 'Enter the table:',
      },
    ], 
    actions: [
      async function(answers) {
        try {
          const { tableName } = answers;
          const data = await fetchDataFromDB(tableName);
          const outputFilePath = `generated/${tableName}.json`;
          await writeFileAsync(outputFilePath, JSON.stringify(data, null, 2));
          return `Generated file with data from database: ${outputFilePath}`;
        } catch (error) {
          console.error('Error generating files:', error);
          throw error;
        }        
      }
    ],
  });
  plop.setGenerator('allData', {
    description: 'Generate all data from database',
    prompts: [], 
    actions: [async () => {
      try {
        const data = await fetchTablesAndData();
        const outputFilePath = 'generated/dataFromDB.json';
        await writeFileAsync(outputFilePath, JSON.stringify(data, null, 2));
        return `Generated file with data from database: ${outputFilePath}`;
      } catch (error) {
        console.error('Error generating files:', error);
        throw error;
      }
    }],
  });
};

