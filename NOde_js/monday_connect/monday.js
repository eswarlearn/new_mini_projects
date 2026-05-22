require('dotenv').config();
const axios = require('axios');

const API_TOKEN = process.env.MONDAY_API_TOKEN;
const BOARD_ID = process.env.YOUR_BOARD_ID;
const API_URL = 'https://api.monday.com/v2';

async function getItemWithUpdates() {
  const query = `
    query {
      boards(ids: [${BOARD_ID}]) {
        items_page {
          items {
            id
            name
            subitems {
              id
              name
              column_values {
                id
                text
                value
              }
            }
            updates {
              id
              body
              created_at
              creator {
                name
              }
              replies {
                id
                body
                created_at
                creator {
                  name
                }
              }
            }
          }
        }
      }
    }
  `;

  try {
    const response = await axios.post(API_URL, { query }, {
      headers: {
        'Authorization': API_TOKEN,
        'Content-Type': 'application/json',
        'API-Version': '2024-01'
      }
    });

    if (response.data.errors) {
      console.log('API Errors:', JSON.stringify(response.data.errors, null, 2));
      return;
    }

    const items = response.data.data.boards[0].items_page.items;

    // Filter only Brinjal or loop all
    items.forEach(item => {
      console.log(`\n=== Task: ${item.name} (ID: ${item.id}) ===`);

      // Subitems
      if (item.subitems && item.subitems.length > 0) {
        console.log(`  Subtasks (${item.subitems.length}):`);
        item.subitems.forEach(sub => {
          console.log(`    - ${sub.name} (ID: ${sub.id})`);
          sub.column_values.forEach(col => {
            if (col.text) console.log(`      [Col] ${col.id}: ${col.text}`);
          });
        });
      }

      // Updates/Comments
      if (item.updates && item.updates.length > 0) {
        console.log(`  Updates (${item.updates.length}):`);
        item.updates.forEach(update => {
          console.log(`    - [${update.created_at}] ${update.creator.name}: ${update.body}`);
          if (update.replies && update.replies.length > 0) {
            update.replies.forEach(reply => {
              console.log(`      Reply [${reply.created_at}] ${reply.creator.name}: ${reply.body}`);
            });
          }
        });
      }
    });

  } catch (error) {
    console.log('Request failed:', error.response?.data || error.message);
  }
}

getItemWithUpdates();