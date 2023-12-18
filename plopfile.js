require('dotenv').config();

const path = require('path');
const handlebars = require('handlebars'); 
const { Client } = require('pg');
const { promisify } = require('util');
const fs = require('fs');
const changeCase = require('lodash');
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

const fetchTablesType = async () => {
  try {
    await client.connect();
    const tableResult = await client.query(
      "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'"
    );
    const tables = tableResult.rows.map(row => row.table_name);
    const allTableData = {};

    for (const table of tables) {
      const result = await client.query(`SELECT column_name, data_type, is_nullable, character_maximum_length FROM information_schema.columns WHERE table_schema = 'public' AND table_name = '${table}'`);
      //mapping 
      const columns = result.rows.map(column => ({
        column_name: column.column_name,
        data_type: column.data_type,
        is_nullable: column.is_nullable === 'YES',
        max_length: column.character_maximum_length,
      }));

      allTableData[table] = columns;
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
        const outputFilePath = 'generated/${tableName}.go';
        await writeFileAsync(outputFilePath, JSON.stringify(data, null, 2));
        return `Generated file with data from database: ${outputFilePath}`;
      } catch (error) {
        console.error('Error generating files:', error);
        throw error;
      }
    }],
  });
  
  /* Compile every file*/ 
  plop.setGenerator('generateGolang', {
    description: 'Generate models from database',
    prompts: [],
    actions: [async () => {
      try {
        const data = await fetchTablesType();
        const outputFolderPath = `generated/models`;
        if (!fs.existsSync(outputFolderPath)) {
          fs.mkdirSync(outputFolderPath, { recursive: true });
        }
        const outputRepoFolderPath = `generated/repository`;
        if (!fs.existsSync(outputRepoFolderPath)) {
          fs.mkdirSync(outputRepoFolderPath, { recursive: true });
        }
        const outputCtrlFolderPath = `generated/controller`;
        if (!fs.existsSync(outputCtrlFolderPath)) {
          fs.mkdirSync(outputCtrlFolderPath, { recursive: true });
        }    
        
        /** Initialize Routes Template  */
        let renderRoute = "";
        const templateRoutesPath = 'templates/goRoutes.hbs'; 
        const templateRoutesContentPath = 'templates/goRoutesContent.hbs'; 
       

        for (const tableName in data) {
          if (Object.hasOwnProperty.call(data, tableName)) {
            const columns = data[tableName];

            const tableData = {
              project_name: 'golangtemplate',
              param_table: tableName,
              table_name: changeCase.upperFirst(changeCase.camelCase(tableName)) ,
                columns: columns.map(column => ({
                column_name: changeCase.upperFirst(changeCase.camelCase(column.column_name)),
                json_column_name: column.column_name,
                data_type: column.data_type,
                max_length: column.character_maximum_length,
                is_nullable: column.is_nullable === 'YES',
              })),      
            };

            /**@@@@@@@@@@@@@@@@@@@@@@@      MODELS      @@@@@@@@@@@@@@@@@@@@@@@@@@@ */
            const templatePath = 'templates/golangModel.hbs';
            const outputFilePath = path.join(outputFolderPath, `${tableName}Model.go`);   

            const templateFileContent = fs.readFileSync(templatePath, 'utf8');
            const compiledTemplate = handlebars.compile(templateFileContent);
            const renderedTemplate = compiledTemplate(tableData);
 
            fs.writeFileSync(outputFilePath, renderedTemplate)
            console.log(`Generated file for ${tableName} at ${outputFilePath}`);

            /**@@@@@@@@@@@@@@@@@@@@@@@      REPOSITORY      @@@@@@@@@@@@@@@@@@@@@@@@@@@ */          
            const templateRepoPath = 'templates/goRepo.hbs';
            const outputRepoFilePath = path.join(outputRepoFolderPath, `${tableName}Repo.go`);   

            const templateRepoFileContent = fs.readFileSync(templateRepoPath, 'utf8');
            const compiledRepoTemplate = handlebars.compile(templateRepoFileContent);
            const renderedRepoTemplate = compiledRepoTemplate(tableData);
 
            fs.writeFileSync(outputRepoFilePath, renderedRepoTemplate)

            /**@@@@@@@@@@@@@@@@@@@@@@@      CONTROLLER      @@@@@@@@@@@@@@@@@@@@@@@@@@@ */          
            const templateCtrlPath = 'templates/goCtrl.hbs';
            const outputCtrlFilePath = path.join(outputCtrlFolderPath, `${tableName}Ctrl.go`);   

            const templateCtrlFileContent = fs.readFileSync(templateCtrlPath, 'utf8');
            const compiledCtrlTemplate = handlebars.compile(templateCtrlFileContent);
            const renderedCtrlTemplate = compiledCtrlTemplate(tableData);
  
            fs.writeFileSync(outputCtrlFilePath, renderedCtrlTemplate)

            /**@@@@@@@@@@@@@@@@@@@@@@@     RENDER ROUTES  CONTENT     @@@@@@@@@@@@@@@@@@@@@@@@@@@ */      
      
            const templateRoutesFileContent = fs.readFileSync(templateRoutesContentPath, 'utf8');
            const compiledRoutesTemplate = handlebars.compile(templateRoutesFileContent);
            const renderedRoutesTemplate = compiledRoutesTemplate(tableData);
            renderRoute += renderedRoutesTemplate;

           
          } 
        }

        //console.log(renderRoute);
        /**@@@@@@@@@@@@@@@@@@@@@@@     MERGE ROUTES CONTENT     @@@@@@@@@@@@@@@@@@@@@@@@@@@ */  
        const routeData = {
          route_content :renderRoute
        } 
       
        const templateRoutesFile = fs.readFileSync(templateRoutesPath,'utf8');
        const compiledRoutesMainTemplate = handlebars.compile(templateRoutesFile);
        const renderedRoutesMainTemplate = compiledRoutesMainTemplate(routeData);

        const outputRoutesFilePath = `generated/main.go`; 
        fs.writeFileSync(outputRoutesFilePath, renderedRoutesMainTemplate )
      

        return 'Files generated successfully';
      } catch (error) {
        console.error('Error generating files:', error);
        throw error;
      }},
    ],
  });
};


handlebars.registerHelper('golangType', function(dataType) {
  if (dataType === 'int' || dataType === 'bigint' || dataType === 'smallint') {
    return 'int64';
  } else if (dataType === 'varchar' || dataType === 'text' || dataType === 'character varying') {
    return 'string';
  } else if (dataType === 'timestamp with time zone' || dataType === 'timestamp'){
    return 'time.Time';
  } else if(dataType === 'boolean'){
    return 'bool';
  } else {
    return 'interface{}';
  }
});

