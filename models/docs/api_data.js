define({ "api": [
  {
    "type": "DELETE",
    "url": "/api/content/:id",
    "title": "Remove Content",
    "description": "<p>Remove Content</p>",
    "group": "Content",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "filename": "./models/content/remove.go",
    "groupTitle": "Content",
    "name": "DeleteApiContentId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/content",
    "title": "Content List",
    "description": "<p>Content List</p>",
    "group": "Content",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "access_level",
            "description": "<p>Content Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Content Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Content Created User ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "pages.id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "pages.active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages.link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages.name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages.title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages.keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pages.descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "pages.json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "pages.access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "pages.created_by",
            "description": "<p>Page Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"name\": \"header intro\",\n\t\t\t\"content\": \"<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>\",\n\t\t\t\"access_level\": null,\n\t\t\t\"json_settings\": [{}],\n\t\t\t\"created_by\": 1,\n\t\t\t\"pages\":\n\t\t\t[\n\t\t\t\t{\n\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\"active\": 1,\n\t\t\t\t\t\"name\": \"Home\",\n\t\t\t\t\t\"link\": \"\",\n\t\t\t\t\t\"title\": \"Rapido Home Page\",\n\t\t\t\t\t\"keywords\": \"Rapido,home,page\",\n\t\t\t\t\t\"descriptions\": null,\n\t\t\t\t\t\"access_level\": 0,\n\t\t\t\t\t\"created_by\": 1\n\t\t\t\t}\n\t\t\t]\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/content/list.go",
    "groupTitle": "Content",
    "name": "GetApiContent",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/content/page",
    "title": "Content On Page",
    "description": "<p>Content On Page</p>",
    "group": "Content",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>PageContent ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "content_id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "location",
            "description": "<p>Location</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "column",
            "description": "<p>Column</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "content",
            "description": "<p>Content Object</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "content.id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "content.access_level",
            "description": "<p>Content Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "content.json_settings",
            "description": "<p>Content Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "content.created_by",
            "description": "<p>Content Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"page_id\": 1,\n\t\t\t\"content_id\": 1,\n\t\t\t\"extension\": \"test\",\n\t\t\t\"location\": \"header\",\n\t\t\t\"column\": 1,\n\t\t\t\"position\": 1,\n\t\t\t\"access_level\": null,\n\t\t\t\"json_settings\": [{}],\n\t\t\t\"content\": {\n\t\t\t    \"id\": 1,\n\t\t\t    \"name\": \"header\",\n\t\t\t    \"content\": \"<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>\",\n\t\t\t    \"access_level\": null,\n\t\t\t    \"created_by\": 1\n\t\t\t}\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "link",
            "description": "<p>Page Link</p>"
          }
        ]
      }
    },
    "filename": "./models/content/get.go",
    "groupTitle": "Content",
    "name": "GetApiContentPage",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/content",
    "title": "Change Content On Page",
    "description": "<p>Change Content On Page</p>",
    "group": "Content",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"page_id\": 1,\n\t\t\t\"content_id\": 1,\n\t\t\t\"extension\": \"test\",\n\t\t\t\"location\": \"header\",\n\t\t\t\"column\": 1,\n\t\t\t\"position\": 1,\n\t\t\t\"access_level\": null,\n\t\t\t\"json_settings\": [{}],\n\t\t\t\"content\": {\n\t\t\t    \"id\": 1,\n\t\t\t    \"name\": \"header\",\n\t\t\t    \"content\": \"<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>\",\n\t\t\t    \"access_level\": null\n\t\t\t}\n\t\t}\n]",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Update page success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "id",
            "description": "<p>PageContent ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "content_id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "location",
            "description": "<p>Location</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "column",
            "description": "<p>Column</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "content",
            "description": "<p>Content Object</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "content.id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "content.name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content.content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "content.access_level",
            "description": "<p>Content Access Level</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "content.json_settings",
            "description": "<p>Content Setting</p>"
          }
        ]
      }
    },
    "filename": "./models/content/add.go",
    "groupTitle": "Content",
    "name": "PostApiContent",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/content/:id",
    "title": "Change Content",
    "description": "<p>Change Content</p>",
    "group": "Content",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"name\": \"Title\",\n \t\"content\": \"<h1>Rapido</h1>\",\n \t\"json_settings\": null,\n \t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Content Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Content Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Content Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"id\": 1,\n \t\"name\": \"Title\",\n \t\"content\": \"<h1>Rapido</h1>\",\n \t\"json_settings\": null,\n \t\"access_level\": 0,\n \t\"created_by\": 1\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Content Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Content Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/content/change.go",
    "groupTitle": "Content",
    "name": "PutApiContentId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/extension",
    "title": "Extension List",
    "description": "<p>Extension List</p>",
    "group": "Extension",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "author",
            "description": "<p>Author Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "content",
            "description": "<p>Content Detail</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created",
            "description": "<p>Created Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>Extension Description</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Unique Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handling Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "layout",
            "description": "<p>Layout Description</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "manifest",
            "description": "<p>Manifest Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "view",
            "description": "<p>View Name</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "menus",
            "description": "<p>Menu Detail</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "version",
            "description": "<p>Extension Version</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "settings",
            "description": "<p>Extension Settings</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "layouts",
            "description": "<p>Extension Layout Detail</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"author\": \"smiffy6969\",\n\t\t\t\"created\": 1394394028,\n\t\t\t\"description\": \"Dynamic contact form.\",\n\t\t\t\"extension\": \"Contact-form\",\n\t\t\t\"handle\": \"razorcms\"\n\t\t\t\"view\": \"contact-form\",\n\t\t\t\"name\": \"Contact Form\",\n\t\t\t\"type\": \"Communication\",\n\t\t\t\"version\": 6,\n\t\t\t\"settings\": [{\n\t\t\t\t\"name\": \"to\",\n\t\t\t\t\"placeholder\": \"Email address to send messages to\",\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"value\": \"\"},\n\t\t{\n\t\t\t\t\"name\": \"layout\",\n\t\t\t\t\"placeholder\": \"Email template\",\n\t\t\t\t\"type\": \"textarea\",\n\t\t\t\t\"value\": \"\"}]",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "size": "\"system\",\"communication\"",
            "optional": true,
            "field": "type",
            "description": "<p>Extension Type</p>"
          }
        ]
      }
    },
    "filename": "./models/extension/list.go",
    "groupTitle": "Extension",
    "name": "GetApiExtension",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/extension",
    "title": "Change Extension Content",
    "description": "<p>Change Extension Content</p>",
    "group": "Extension",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"id\": 1,\n\t\t\"extension\": \"basic-blue-side\",\n\t\t\"type\": \"communication\",\n\t\t\"handle\": \"rapidohandle\",\n\t\t\"json_settings\": null,\n\t\t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Extension ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handle Method</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Extension Default Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Extension Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "user_id",
            "description": "<p>Extension Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"id\": 1,\n\t\t\"extension\": \"contact-form\",\n\t\t\"type\": \"communication\",\n\t\t\"handle\": \"rapidohandle\",\n\t\t\"json_settings\": null,\n\t\t\"user_id\": 1,\n\t\t\"access_level\": 0\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "id",
            "description": "<p>Extension ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Type</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handle Method</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/extension/add.go",
    "groupTitle": "Extension",
    "name": "PostApiExtension",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/extension/repository",
    "title": "Detail Extension Repository",
    "description": "<p>Detail Extension Repository</p>",
    "group": "Extension",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"type\": \"communication\",\n\t\t\"handle\": \"rapidohandle\",\n\t\t\"extension\": \"contact-form\",\n\t\t\"manifest\": \"\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "author",
            "description": "<p>Author Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "content",
            "description": "<p>Content Detail</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created",
            "description": "<p>Created Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>Extension Description</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Unique Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handling Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "layout",
            "description": "<p>Layout Description</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "manifest",
            "description": "<p>Manifest Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "view",
            "description": "<p>View Name</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "menus",
            "description": "<p>Menu Detail</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "version",
            "description": "<p>Extension Version</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "settings",
            "description": "<p>Extension Settings</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "layouts",
            "description": "<p>Extension Layout Detail</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\t\"author\": \"smiffy6969\",\n\t\t\t\"created\": 1394394028,\n\t\t\t\"description\": \"Dynamic contact form.\",\n\t\t\t\"extension\": \"Contact-form\",\n\t\t\t\"handle\": \"razorcms\"\n\t\t\t\"view\": \"contact-form\",\n\t\t\t\"name\": \"Contact Form\",\n\t\t\t\"type\": \"Communication\",\n\t\t\t\"version\": 6,\n\t\t\t\"settings\": [{\n\t\t\t\t\"name\": \"to\",\n\t\t\t\t\"placeholder\": \"Email address to send messages to\",\n\t\t\t\t\"type\": \"text\",\n\t\t\t\t\"value\": \"\"},\n\t\t{\n\t\t\t\t\"name\": \"layout\",\n\t\t\t\t\"placeholder\": \"Email template\",\n\t\t\t\t\"type\": \"textarea\",\n\t\t\t\t\"value\": \"\"}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Category</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handler</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "manifest",
            "description": "<p>Manifest Name</p>"
          }
        ]
      }
    },
    "filename": "./models/extension/repository.go",
    "groupTitle": "Extension",
    "name": "PostApiExtensionRepository",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/plugins",
    "title": "Call Extension",
    "description": "<p>Call Extension</p>",
    "group": "Extension",
    "version": "0.1.0",
    "permission": [
      {
        "name": "none"
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"extension\": \"contact-form\",\n\t\t\"type\": \"communication\",\n\t\t\"handle\": \"rapidohandle\",\n\t\t\"data\": {}\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Extension Type</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "handle",
            "description": "<p>Extension Handle Method</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>Plugin Param Data</p>"
          }
        ]
      }
    },
    "filename": "./models/plugins/plugins.go",
    "groupTitle": "Extension",
    "name": "PostApiPlugins",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/api/file/media",
    "title": "Remove Media",
    "description": "<p>Remove Media</p>",
    "group": "File",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Media Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Media Type</p>"
          }
        ]
      }
    },
    "filename": "./models/file/remove_media.go",
    "groupTitle": "File",
    "name": "DeleteApiFileMedia",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/file/backup",
    "title": "Backup Project",
    "description": "<p>Backup Project</p>",
    "group": "File",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"message\": \"upgrade_backup.zip\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/file/backup.go",
    "groupTitle": "File",
    "name": "GetApiFileBackup",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/file/backup_list",
    "title": "Backup List",
    "description": "<p>Backup List</p>",
    "group": "File",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Image Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Timestamp",
            "optional": false,
            "field": "created_on",
            "description": "<p>Data Modified Time</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"name\": \"upgrade_backup.zip\",\n\t\t\t\"created_on\": \"2018-01-02 11:20:31\"\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/file/backup_list.go",
    "groupTitle": "File",
    "name": "GetApiFileBackup_list",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/file/media",
    "title": "Media List",
    "description": "<p>Media List</p>",
    "group": "File",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Media Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>Media Type</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "url",
            "description": "<p>Media URL</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"name\": \"image_a.jpeg\",\n\t\t\t\"type\": \"image\",\n\t\t\t\"url\": \"http://rapidohandle.com/storage/file/image/image_a.jpeg\"\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/file/media_list.go",
    "groupTitle": "File",
    "name": "GetApiFileMedia",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/file/media",
    "title": "Add Media",
    "description": "<p>Add Media</p>",
    "group": "File",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>File name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "url",
            "description": "<p>Media URL</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"name\": \"image_a.jpeg\",\n\t\t\t\"url\": \"http://rapidohandle.com/storage/file/image/image_a.jpeg\"\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "[]Blob",
            "optional": false,
            "field": "files",
            "description": "<p>Media File. Support only (gif,png,jpg,ogg,mp3,flac,avi,mpg,mkv,webm,pdf,txt,odt,ott,ods,ots,odp,otp,odg,otg)</p>"
          }
        ]
      }
    },
    "filename": "./models/file/add_media.go",
    "groupTitle": "File",
    "name": "PostApiFileMedia",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/menu",
    "title": "Menu List",
    "description": "<p>Menu List</p>",
    "group": "Menu",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Menu Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Access Level Menu</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"name\": \"header\",\n\t\t\t\"access_level\": null\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/menu/list.go",
    "groupTitle": "Menu",
    "name": "GetApiMenu",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/menu_item",
    "title": "Menu Item",
    "description": "<p>Menu Item</p>",
    "group": "Menu",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Menu Name</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "item",
            "description": "<p>Menu Item</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.id",
            "description": "<p>Menu Item ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.menu_id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.level",
            "description": "<p>Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.link_label",
            "description": "<p>Link Label</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.link_url",
            "description": "<p>Link URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.link_target",
            "description": "<p>Link Target</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "item.page",
            "description": "<p>Item Page</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.page.id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.page.active",
            "description": "<p>Page Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.page.name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.page.link",
            "description": "<p>Page Link</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "item.sub_menu",
            "description": "<p>Sub Menu Item</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.id",
            "description": "<p>Menu Item ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.menu_id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.level",
            "description": "<p>Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.sub_menu.link_label",
            "description": "<p>Link Label</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.sub_menu.link_url",
            "description": "<p>Link URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.sub_menu.link_target",
            "description": "<p>Link Target</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "item.sub_menu.page",
            "description": "<p>Item Page</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.page.id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "item.sub_menu.page.active",
            "description": "<p>Page Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.sub_menu.page.name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "item.sub_menu.page.link",
            "description": "<p>Page Link</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"name\": \"header\",\n\t\t\t\"item\":\n\t\t\t[\n\t\t\t\t{\n\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\"menu_id\": 1,\n\t\t\t\t\t\"position\": 1,\n\t\t\t\t\t\"level\": 1,\n\t\t\t\t\t\"page_id\": 1,\n\t\t\t\t\t\"link_label\": \"\",\n\t\t\t\t\t\"link_url\": \"\",\n\t\t\t\t\t\"link_target\": \"\",\n\t\t\t\t\t\"page\":\n\t\t\t\t\t{\n\t\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\t\"active\": 1,\n\t\t\t\t\t\t\"name\": \"Home\",\n\t\t\t\t\t\t\"link\": \"/index\"\n\t\t\t\t\t},\n\t\t\t\t\t\"sub_menu\":\n\t\t\t\t\t[\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"id\": 4,\n\t\t\t\t\t\t\t\"menu_id\": 1,\n\t\t\t\t\t\t\t\"position\": 1,\n\t\t\t\t\t\t\t\"level\": 2,\n\t\t\t\t\t\t\t\"page_id\": 1,\n\t\t\t\t\t\t\t\"link_label\": \"\",\n\t\t\t\t\t\t\t\"link_url\": \"\",\n\t\t\t\t\t\t\t\"link_target\": \"\",\n\t\t\t\t\t\t\t\"page\":\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\t\t\t\"active\": 1,\n\t\t\t\t\t\t\t\t\"name\": \"Home\",\n\t\t\t\t\t\t\t\t\"link\": \"/index\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/menu/get.go",
    "groupTitle": "Menu",
    "name": "GetApiMenu_item",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/menu",
    "title": "Change Menu Item",
    "description": "<p>Change Menu Item</p>",
    "group": "Menu",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"name\": \"header\",\n\t\t\t\"item\":\n\t\t\t[\n\t\t\t\t{\n\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\"menu_id\": 1,\n\t\t\t\t\t\"position\": 1,\n\t\t\t\t\t\"level\": 1,\n\t\t\t\t\t\"page_id\": 1,\n\t\t\t\t\t\"link_label\": \"\",\n\t\t\t\t\t\"link_url\": \"\",\n\t\t\t\t\t\"link_target\": \"\",\n\t\t\t\t\t\"page\":\n\t\t\t\t\t{\n\t\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\t\"active\": 1,\n\t\t\t\t\t\t\"name\": \"Home\",\n\t\t\t\t\t\t\"link\": \"/index\"\n\t\t\t\t\t},\n\t\t\t\t\t\"sub_menu\":\n\t\t\t\t\t[\n\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\"id\": 4,\n\t\t\t\t\t\t\t\"menu_id\": 1,\n\t\t\t\t\t\t\t\"position\": 1,\n\t\t\t\t\t\t\t\"level\": 2,\n\t\t\t\t\t\t\t\"page_id\": 1,\n\t\t\t\t\t\t\t\"link_label\": \"\",\n\t\t\t\t\t\t\t\"link_url\": \"\",\n\t\t\t\t\t\t\t\"link_target\": \"\",\n\t\t\t\t\t\t\t\"page\":\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\t\t\t\"active\": 1,\n\t\t\t\t\t\t\t\t\"name\": \"Home\",\n\t\t\t\t\t\t\t\t\"link\": \"/index\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t]\n\t\t}\n]",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Update menu success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Menu Name</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": false,
            "field": "item",
            "description": "<p>Menu Item</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.id",
            "description": "<p>Menu Item ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.menu_id",
            "description": "<p>Menu ID (1 for header, 2 for footer, 3 for main)</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.level",
            "description": "<p>Level (0 for root, 1 for child)</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "item.link_label",
            "description": "<p>Link Label</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "item.link_url",
            "description": "<p>Link URL</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "item.link_target",
            "description": "<p>Link Target</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "item.page",
            "description": "<p>Item Page</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.page.id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "item.page.active",
            "description": "<p>Page Active Status</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "item.page.name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "item.page.link",
            "description": "<p>Page Link</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "item.sub_menu",
            "description": "<p>Sub Menu Item</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.id",
            "description": "<p>Menu Item ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.menu_id",
            "description": "<p>Menu ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.level",
            "description": "<p>Level</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "item.sub_menu.link_label",
            "description": "<p>Link Label</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "item.sub_menu.link_url",
            "description": "<p>Link URL</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "item.sub_menu.link_target",
            "description": "<p>Link Target</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": true,
            "field": "item.sub_menu.page",
            "description": "<p>Item Page</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.page.id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "item.sub_menu.page.active",
            "description": "<p>Page Active Status</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "item.sub_menu.page.name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "item.sub_menu.page.link",
            "description": "<p>Page Link</p>"
          }
        ]
      }
    },
    "filename": "./models/menu/change.go",
    "groupTitle": "Menu",
    "name": "PostApiMenu",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/api/page/:archive_id/archive",
    "title": "Remove Archive Page",
    "description": "<p>Remove Archive Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/remove_archive.go",
    "groupTitle": "Page",
    "name": "DeleteApiPageArchive_idArchive",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/api/page/:id",
    "title": "Remove Page",
    "description": "<p>Remove Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/remove.go",
    "groupTitle": "Page",
    "name": "DeleteApiPageId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/api/page/:id/collaborator/:email",
    "title": "Remove Page Collaborators",
    "description": "<p>Remove Page Collaborators</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/remove_collaborator.go",
    "groupTitle": "Page",
    "name": "DeleteApiPageIdCollaboratorEmail",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/archive/page",
    "title": "Page Archive List",
    "description": "<p>Page Archive List</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "data",
            "description": "<p>Page Version Data</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 2,\n\t\t\t\"link\": \"test\",\n\t\t\t\"name\": \"Home\",\n\t\t\t\"title\": \"Rapido Home Page\",\n\t\t\t\"data\": [{}]\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/list_archive.go",
    "groupTitle": "Page",
    "name": "GetApiArchivePage",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/page",
    "title": "Page List",
    "description": "<p>Page List</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Page Created User ID</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "collaborators",
            "description": "<p>Email Of Collaborators</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 2,\n\t\t\t\"active\": 1,\n\t\t\t\"link\": \"test\",\n\t\t\t\"name\": \"Home\",\n\t\t\t\"title\": \"Rapido Home Page\",\n\t\t\t\"keywords\": \"Rapido,home,page\",\n\t\t\t\"descriptions\": null,\n\t\t\t\"json_settings\": null,\n\t\t\t\"access_level\": 0,\n\t\t\t\"created_by\": 1,\n\t\t\t\"collaborators\": [\"email@example.com\"]\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/list.go",
    "groupTitle": "Page",
    "name": "GetApiPage",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/page/:id",
    "title": "Page Details",
    "description": "<p>Page Details</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "collaborators",
            "description": "<p>Email Of Collaborators</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"id\": 2,\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0,\n\t\t\"collaborators\": [\"email@example.com\"]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/get.go",
    "groupTitle": "Page",
    "name": "GetApiPageId",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/page/:id/permission",
    "title": "Check Page Permission",
    "description": "<p>Check Page Permission</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Access granted\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/permission.go",
    "groupTitle": "Page",
    "name": "GetApiPageIdPermission",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/page/:page_id/version",
    "title": "Page Version List",
    "description": "<p>Page Version List</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Unique ID Version</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "version",
            "description": "<p>Page Version Number</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "data",
            "description": "<p>Page Data</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.id",
            "description": "<p>PageContent ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.page_id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.content_id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.extension",
            "description": "<p>Extension Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.location",
            "description": "<p>Location</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.column",
            "description": "<p>Column</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.position",
            "description": "<p>Position</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "data.json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data.content",
            "description": "<p>Content Object</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.content.id",
            "description": "<p>Content ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.content.name",
            "description": "<p>Content Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.content.content",
            "description": "<p>Content</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "data.content.access_level",
            "description": "<p>Content Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "data.content.json_settings",
            "description": "<p>Content Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Version Created By User ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_on",
            "description": "<p>Version Created Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "current",
            "description": "<p>Current Version Flag</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name Of User</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"page_id\": 1,\n\t\t\t\"version\": 1,\n\t\t\t\"data\":\n\t\t\t[\n\t\t\t\t{\n\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\"page_id\": 1,\n\t\t\t\t\t\"content_id\": 1,\n\t\t\t\t\t\"extension\": \"test\",\n\t\t\t\t\t\"location\": \"header\",\n\t\t\t\t\t\"column\": 1,\n\t\t\t\t\t\"position\": 1,\n\t\t\t\t\t\"access_level\": null,\n\t\t\t\t\t\"json_settings\": [{}],\n\t\t\t\t\t\"content\":\n\t\t\t\t\t{\n\t\t\t\t\t\t\"id\": 1,\n\t\t\t\t\t\t\"name\": \"header\",\n\t\t\t\t\t\t\"content\": \"<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>\",\n\t\t\t\t\t\t\"access_level\": null\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t],\n\t\t\t\"created_by\": 1,\n\t\t\t\"created_on\": \"1517677505\",\n\t\t\t\"current\": 1,\n\t\t\t\"name\": \"admin\"\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/list_version.go",
    "groupTitle": "Page",
    "name": "GetApiPagePage_idVersion",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/page",
    "title": "Add New Page",
    "description": "<p>Add New Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"active\": 1,\n\t\t\"link\": \"test\",\n\t\t\"name\": \"Home\",\n\t\t\"title\": \"Rapido Home Page\",\n\t\t\"keywords\": \"Rapido,home,page\",\n\t\t\"descriptions\": null,\n\t\t\"json_settings\": null,\n\t\t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Page Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created\n{\n\t\t\"id\": 1,\n\t\t\"active\": 1,\n\t\t\"link\": \"test\",\n\t\t\"name\": \"Home\",\n\t\t\"title\": \"Rapido Home Page\",\n\t\t\"keywords\": \"Rapido,home,page\",\n\t\t\"descriptions\": null,\n\t\t\"json_settings\": null,\n\t\t\"access_level\": 0,\n\t\t\"created_by\": 1\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/page/add.go",
    "groupTitle": "Page",
    "name": "PostApiPage",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/page/:id/collaborator",
    "title": "Add Page Collaborator",
    "description": "<p>Add Page Collaborator</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"email\": \"sample@email.com\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created\n{\n\t\t\"message\": \"Add collaborator success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>Collaborator email</p>"
          }
        ]
      }
    },
    "filename": "./models/page/add_collaborator.go",
    "groupTitle": "Page",
    "name": "PostApiPageIdCollaborator",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/page/:id/copy",
    "title": "Copy Page",
    "description": "<p>Copy Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Page Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created\n{\n \t\"id\": 2,\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0,\n \t\"created_by\": 1\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/page/copy.go",
    "groupTitle": "Page",
    "name": "PostApiPageIdCopy",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/page/:id/share",
    "title": "Share Page",
    "description": "<p>Share Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"email\": \"sample@email.com\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created\n{\n\t\t\"message\": \"Share page success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>Collaborator email</p>"
          }
        ]
      }
    },
    "filename": "./models/page/share.go",
    "groupTitle": "Page",
    "name": "PostApiPageIdShare",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/page/:archive_id/restore",
    "title": "Restore Page",
    "description": "<p>Restore Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Restore page success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/restore.go",
    "groupTitle": "Page",
    "name": "PutApiPageArchive_idRestore",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/page/:id",
    "title": "Change Page",
    "description": "<p>Change Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Page Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"id\": 1,\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0,\n \t\"created_by\": 1\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": true,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/page/change.go",
    "groupTitle": "Page",
    "name": "PutApiPageId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/page/:id/access_level",
    "title": "Change Page Access Level",
    "description": "<p>Change Page Access Level</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"access_level\": 0\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Status</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "link",
            "description": "<p>URL Link</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Page Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>Page Title</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keywords",
            "description": "<p>Page Keywords</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "descriptions",
            "description": "<p>Page Descriptions</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "json_settings",
            "description": "<p>Page Setting</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>Page Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "created_by",
            "description": "<p>Page Created User ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"id\": 1,\n \t\"active\": 1,\n \t\"link\": \"test\",\n \t\"name\": \"Home\",\n \t\"title\": \"Rapido Home Page\",\n \t\"keywords\": \"Rapido,home,page\",\n \t\"descriptions\": null,\n \t\"json_settings\": null,\n \t\"access_level\": 0,\n \t\"created_by\": 1\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": ""
          }
        ]
      }
    },
    "filename": "./models/page/change_al.go",
    "groupTitle": "Page",
    "name": "PutApiPageIdAccess_level",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/page/:version_id/revert",
    "title": "Revert Page",
    "description": "<p>Revert Page</p>",
    "group": "Page",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Revert page success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/page/revert.go",
    "groupTitle": "Page",
    "name": "PutApiPageVersion_idRevert",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/repository/:type",
    "title": "Repository List",
    "description": "<p>Repository List</p>",
    "group": "Repository",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "extensions",
            "description": "<p>Extension List</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "categories",
            "description": "<p>Category List</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Object",
            "optional": false,
            "field": "handles",
            "description": "<p>Handle List</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n \t\"extensions\": [{}],\n \t\"categories\": [{}],\n \t\"handles\": [{}]\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "size": "\"extension\",\"category\",\"handle\"",
            "optional": false,
            "field": "type",
            "description": "<p>Type of repository</p>"
          }
        ]
      }
    },
    "filename": "./models/repository/list.go",
    "groupTitle": "Repository",
    "name": "GetApiRepositoryType",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/logo",
    "title": "Get Logo And Favicon URL",
    "description": "<p>Get Logo And Favicon URL</p>",
    "group": "Setting",
    "version": "0.1.0",
    "permission": [
      {
        "name": "none"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "favicon_url",
            "description": "<p>Favicon URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "logo_url",
            "description": "<p>Logo URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name Of Website</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "max_size_upload",
            "description": "<p>Maximum Limit For Uploading File in MB</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"favicon_url\": \"\",\n\t\t\"logo_url\": \"\",\n\t\t\"name\": \"my website\",\n\t\t\"max_size_upload\": 50\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/setting/logo.go",
    "groupTitle": "Setting",
    "name": "GetApiLogo",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/setting",
    "title": "Setting Parameter",
    "description": "<p>Setting Parameter</p>",
    "group": "Setting",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_host",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "smtp_port",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_user",
            "description": "<p>SMTP User For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_password",
            "description": "<p>SMTP Password For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_sender",
            "description": "<p>Email Sender For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activation_email",
            "description": "<p>Activation Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activated_email",
            "description": "<p>Activated Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "forget_password_email",
            "description": "<p>Forget Password Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "page_collaborator",
            "description": "<p>Page Collaborator Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "share_page_email",
            "description": "<p>Share Page Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Site Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "home_page",
            "description": "<p>Home Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "banned_attemps",
            "description": "<p>Login Failure Attemps Before Banned</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "lock_time",
            "description": "<p>Seconds To Lockout For After Above Failures Detected</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_attemps",
            "description": "<p>Attempts Are Allowed Before Lockout</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_expired",
            "description": "<p>Amount Of Seconds Of Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "secret_key",
            "description": "<p>Secret Key For Generate Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "max_size_upload",
            "description": "<p>Maximum Limit For Uploading File in MB</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "favicon_url",
            "description": "<p>Favicon URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "logo_url",
            "description": "<p>Logo URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "remote_check_url",
            "description": "<p>URL For Checking System Version</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "system_upgrade_url",
            "description": "<p>URL For Get Upgrade Package</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension_check_url",
            "description": "<p>URL For Checking Extension List</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"smtp_host\": \"-\",\n\t\t\"smtp_port\": 1,\n\t\t\"smtp_user\": \"-\",\n\t\t\"smtp_password\": \"-\",\n\t\t\"email_sender\": \"-\",\n\t\t\"activation_email\": \"<html><head><title>Activate Account</title></head><body></body></html>\",\n\t\t\"activated_email\": \"<html><head><title>Account Activated</title></head><body></body></html>\",\n\t\t\"forget_password_email\": \"<html><head><title>Password Reset</title></head><body></body></html>\",\n\t\t\"page_collaborator\": \"<html><head><title>Page Collaborator</title></head><body></body></html>\",\n\t\t\"share_page_email\": \"<html><head><title>Share Page</title></head><body></body></html>\",\n\t\t\"name\": \"My Site\",\n\t\t\"home_page\": 1,\n\t\t\"banned_attemps\": 250,\n\t\t\"lock_time\": 1,\n\t\t\"access_attemps\": 5,\n\t\t\"access_expired\": 86400,\n\t\t\"secret_key\": \"HRqchYI51YsPEA6E9H0ITwBM4B5F5Nnf\",\n\t\t\"max_size_upload\": 50,\n\t\t\"favicon_url\": \"\",\n\t\t\"logo_url\": \"\",\n\t\t\"remote_check_url\": \"\",\n\t\t\"system_upgrade_url\": \"\",\n\t\t\"extension_check_url\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/setting/get.go",
    "groupTitle": "Setting",
    "name": "GetApiSetting",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/setting",
    "title": "Change Setting",
    "description": "<p>Change Setting</p>",
    "group": "Setting",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"show_logo\": true,\n\t\t\"show_menus\": false,\n\t\t\"show_locales_menu\": true,\n\t\t\"smtp_host\": \"-\",\n\t\t\"smtp_port\": 1,\n\t\t\"smtp_user\": \"-\",\n\t\t\"smtp_password\": \"-\",\n\t\t\"email_sender\": \"-\",\n\t\t\"activation_email\": \"<html><head><title>Activate Account</title></head><body></body></html>\",\n\t\t\"forget_password_email\": \"<html><head><title>Password Reset</title></head><body></body></html>\",\n\t\t\"name\": \"My Site\",\n\t\t\"home_page\": 1,\n\t\t\"banned_attemps\": 250,\n\t\t\"lock_time\": 1,\n\t\t\"access_attemps\": 5,\n\t\t\"access_expired\": 86400,\n\t\t\"secret_key\": \"HRqchYI51YsPEA6E9H0ITwBM4B5F5Nnf\",\n\t\t\"max_size_upload\": 50,\n\t\t\"favicon_url\": \"\",\n\t\t\"logo_url\": \"\",\n\t\t\"remote_check_url\": \"\",\n\t\t\"system_upgrade_url\": \"\",\n\t\t\"extension_check_url\": \"\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_host",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "smtp_port",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_user",
            "description": "<p>SMTP User For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "smtp_password",
            "description": "<p>SMTP Password For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_sender",
            "description": "<p>Email Sender For Sending Email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activation_email",
            "description": "<p>Activation Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "forget_password_email",
            "description": "<p>Forget Password Email Template</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Site Name</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "home_page",
            "description": "<p>Home Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "banned_attemps",
            "description": "<p>Login Failure Attemps Before Banned</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "lock_time",
            "description": "<p>Seconds To Lockout For After Above Failures Detected</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_attemps",
            "description": "<p>Attempts Are Allowed Before Lockout</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_expired",
            "description": "<p>Amount Of Seconds Of Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "secret_key",
            "description": "<p>Secret Key For Generate Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "max_size_upload",
            "description": "<p>Maximum Limit For Uploading File in MB</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "favicon_url",
            "description": "<p>Favicon URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "logo_url",
            "description": "<p>Logo URL</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "remote_check_url",
            "description": "<p>URL For Checking System Version</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "system_upgrade_url",
            "description": "<p>URL For Get Upgrade Package</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "extension_check_url",
            "description": "<p>URL For Checking Extension List</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"smtp_host\": \"-\",\n\t\t\"smtp_port\": 1,\n\t\t\"smtp_user\": \"-\",\n\t\t\"smtp_password\": \"-\",\n\t\t\"email_sender\": \"-\",\n\t\t\"activation_email\": \"<html><head><title>Activate Account</title></head><body></body></html>\",\n\t\t\"activated_email\": \"<html><head><title>Account Activated</title></head><body></body></html>\",\n\t\t\"forget_password_email\": \"<html><head><title>Password Reset</title></head><body></body></html>\",\n\t\t\"page_collaborator\": \"<html><head><title>Page Collaborator</title></head><body></body></html>\",\n\t\t\"share_page_email\": \"<html><head><title>Share Page</title></head><body></body></html>\",\n\t\t\"name\": \"My Site\",\n\t\t\"home_page\": 1,\n\t\t\"banned_attemps\": 250,\n\t\t\"lock_time\": 1,\n\t\t\"access_attemps\": 5,\n\t\t\"access_expired\": 86400,\n\t\t\"secret_key\": \"HRqchYI51YsPEA6E9H0ITwBM4B5F5Nnf\",\n\t\t\"max_size_upload\": 50,\n\t\t\"favicon_url\": \"\",\n\t\t\"logo_url\": \"\",\n\t\t\"remote_check_url\": \"\",\n\t\t\"extension_check_url\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "smtp_host",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "smtp_port",
            "description": "<p>SMTP Host For Sending Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "smtp_user",
            "description": "<p>SMTP User For Sending Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "smtp_password",
            "description": "<p>SMTP Password For Sending Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "email_sender",
            "description": "<p>Email Sender For Sending Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "activation_email",
            "description": "<p>Activation Email Template</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "activated_email",
            "description": "<p>Activated Email Template</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "forget_password_email",
            "description": "<p>Forget Password Email Template</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "page_collaborator",
            "description": "<p>Page Collaborator Email Template</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "share_page_email",
            "description": "<p>Share Page Email Template</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "name",
            "description": "<p>Site Name</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "home_page",
            "description": "<p>Home Page ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "banned_attemps",
            "description": "<p>Login Failure Attemps Before Banned</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "lock_time",
            "description": "<p>Seconds To Lockout For After Above Failures Detected</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_attemps",
            "description": "<p>Attempts Are Allowed Before Lockout</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "access_expired",
            "description": "<p>Amount Of Seconds Of Token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "secret_key",
            "description": "<p>Secret Key For Generate Token</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": true,
            "field": "max_size_upload",
            "description": "<p>Maximum Limit For Uploading File in MB</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "favicon_url",
            "description": "<p>Favicon URL</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "logo_url",
            "description": "<p>Logo URL</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "remote_check_url",
            "description": "<p>URL For Checking System Version</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "system_upgrade_url",
            "description": "<p>URL For Get Upgrade Package</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "extension_check_url",
            "description": "<p>URL For Checking Extension List</p>"
          }
        ]
      }
    },
    "filename": "./models/setting/change.go",
    "groupTitle": "Setting",
    "name": "PutApiSetting",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/system",
    "title": "System Status",
    "description": "<p>System Status</p>",
    "group": "System",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>System ID</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "version",
            "description": "<p>Version Number</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "milestone",
            "description": "<p>Milestone Number</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "release",
            "description": "<p>Release Number</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"id\": 1,\n\t\t\"version\": 3,\n\t\t\"milestone\": 4,\n\t\t\"release\": 6\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/system/get.go",
    "groupTitle": "System",
    "name": "GetApiSystem",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/system/package",
    "title": "Download System Package",
    "description": "<p>Download System Package</p>",
    "group": "System",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"File save to path\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/system/package.go",
    "groupTitle": "System",
    "name": "GetApiSystemPackage",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/system/upgrade",
    "title": "Upgrade System",
    "description": "<p>Upgrade System</p>",
    "group": "System",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Upgrade version success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/system/upgrade.go",
    "groupTitle": "System",
    "name": "GetApiSystemUpgrade",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/system/revert",
    "title": "Revert System",
    "description": "<p>Revert Status</p>",
    "group": "System",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Success Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"Revert version to upgrade_backup_1520491283.zip success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "filename",
            "description": "<p>Target File To Revert</p>"
          }
        ]
      }
    },
    "filename": "./models/system/revert.go",
    "groupTitle": "System",
    "name": "PostApiSystemRevert",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/version",
    "title": "Get Remote Version",
    "description": "<p>Get Remote Version</p>",
    "group": "Tools",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "version",
            "description": "<p>Version Number</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "milestone",
            "description": "<p>Milestone Number</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "release",
            "description": "<p>Release Number</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"version\": 3,\n\t\t\"milestone\": 4,\n\t\t\"release\": 6\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/tools/version.go",
    "groupTitle": "Tools",
    "name": "GetApiVersion",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/api/user/:id",
    "title": "Remove User",
    "description": "<p>Remove User</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No Content",
          "type": "json"
        }
      ]
    },
    "filename": "./models/user/remove.go",
    "groupTitle": "User",
    "name": "DeleteApiUserId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/activate_user",
    "title": "Activated User",
    "description": "<p>Activated User</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"id\": 1,\n \t\"active_token\": \"abcdefg123456asdfghjbcd\"\n}",
        "type": "form"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>User ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "active_token",
            "description": "<p>Active Token</p>"
          }
        ]
      }
    },
    "filename": "./models/user/activate.go",
    "groupTitle": "User",
    "name": "GetApiActivate_user",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/check_token",
    "title": "Check Token Access",
    "description": "<p>Check Token Access</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "iat",
            "description": "<p>Time Created Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "exp",
            "description": "<p>Time Expired Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "iss",
            "description": "<p>Issuer Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "sub",
            "description": "<p>Subject Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>User ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_logged_in",
            "description": "<p>User Last Login</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"iat\": 12345678,\n\t\t\"exp\": 12345923,\n\t\t\"iss\": \"rapidohandle.com\",\n\t\t\"sub\": \"access granted for 24 hour(s)\",\n\t\t\"id\": 1,\n\t\t\"name\": \"admin\",\n\t\t\"email_address\": \"admin@admin\",\n\t\t\"last_logged_in\": 12345678,\n\t\t\"access_level\": 10\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./models/user/check_token.go",
    "groupTitle": "User",
    "name": "GetApiCheck_token",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/api/user",
    "title": "User List",
    "description": "<p>User List</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Hash Password</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Flag</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_logged_in",
            "description": "<p>Last Login Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_accessed",
            "description": "<p>Last Access Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "failed_attempts",
            "description": "<p>Failed Login Attempts</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "lock_until",
            "description": "<p>Lock Until Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ip_address",
            "description": "<p>IP Address</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activate_token",
            "description": "<p>Activate Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "reminder_token",
            "description": "<p>Reminder Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "reminder_time",
            "description": "<p>Reminder Timestamp</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n[\n\t\t{\n\t\t\t\"id\": 1,\n\t\t\t\"password\": \"ab23fecd32efw34efghe4g5tvij\"\n\t\t\t\"active\": 1,\n\t\t\t\"name\": \"user_test01\",\n\t\t\t\"email_address\": \"test@mail.com\",\n\t\t\t\"last_logged_in\": null,\n\t\t\t\"access_level\": 1,\n\t\t\t\"last_accessed\": null,\n\t\t\t\"failed_attempts\": null,\n\t\t\t\"lock_until\": null,\n\t\t\t\"ip_address\": null,\n\t\t\t\"activate_token\": null,\n\t\t\t\"reminder_token\": null,\n\t\t\t\"reminder_time\": null\n\t\t}\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./models/user/list.go",
    "groupTitle": "User",
    "name": "GetApiUser",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/forget_password",
    "title": "Forget Password",
    "description": "<p>Forget Password</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"email_address\": \"test@mail.com\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          }
        ]
      }
    },
    "filename": "./models/user/forget_password.go",
    "groupTitle": "User",
    "name": "PostApiForget_password",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/login",
    "title": "Login",
    "description": "<p>Login</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"username\": \"user@mail.com\",\n\t\t\"password\": \"abcdefghij\",\n\t\t\"remember\": true\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>Access Token</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 202 Accepted\n{\n\t\t\"token\": \"ab23fecd32efw34efghe4g5tvij.ajfsdhods8sd2323.dsjfhsd9fbwgvwf27fevwcx\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Password</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "remember",
            "description": "<p>Remember Flag</p>"
          }
        ]
      }
    },
    "filename": "./models/user/login.go",
    "groupTitle": "User",
    "name": "PostApiLogin",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/new_password",
    "title": "Renew Password",
    "description": "<p>Renew Password</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n \t\"id\": 1,\n \t\"reminder_token\": \"abcdefg123456asdfghjbcd\",\n \t\"new_password\": \"new_pass\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>User ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "reminder_token",
            "description": "<p>Reminder Token</p>"
          }
        ]
      }
    },
    "filename": "./models/user/renew_password.go",
    "groupTitle": "User",
    "name": "PostApiNew_password",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/register",
    "title": "Register User",
    "description": "<p>Register User</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"password\": \"abcdefghij\"\n\t\t\"name\": \"user_test01\",\n\t\t\"email_address\": \"test@mail.com\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"registration success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          }
        ]
      }
    },
    "filename": "./models/user/register.go",
    "groupTitle": "User",
    "name": "PostApiRegister",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "POST",
    "url": "/api/user",
    "title": "Add New User",
    "description": "<p>Add New User</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Admin",
        "title": "Admin user access only",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"password\": \"abcdefghij\"\n\t\t\"name\": \"user_test01\",\n\t\t\"email_address\": \"test@mail.com\",\n\t\t\"access_level\": 1\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Hash Password</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Flag</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_logged_in",
            "description": "<p>Last Login Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_accessed",
            "description": "<p>Last Access Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "failed_attempts",
            "description": "<p>Failed Login Attempts</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "lock_until",
            "description": "<p>Lock Until Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ip_address",
            "description": "<p>IP Address</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activate_token",
            "description": "<p>Activate Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "reminder_token",
            "description": "<p>Reminder Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "reminder_time",
            "description": "<p>Reminder Timestamp</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created\n{\n\t\t\"id\": 1,\n\t\t\"password\": \"ab23fecd32efw34efghe4g5tvij\"\n\t\t\"active\": 1,\n\t\t\"name\": \"user_test01\",\n\t\t\"email_address\": \"test@mail.com\",\n\t\t\"last_logged_in\": null,\n\t\t\"access_level\": 1,\n\t\t\"last_accessed\": null,\n\t\t\"failed_attempts\": null,\n\t\t\"lock_until\": null,\n\t\t\"ip_address\": null,\n\t\t\"activate_token\": null,\n\t\t\"reminder_token\": null,\n\t\t\"reminder_time\": null\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/user/add.go",
    "groupTitle": "User",
    "name": "PostApiUser",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/change_password",
    "title": "Change Password",
    "description": "<p>Change Password</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Public",
        "title": "Public access",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"password\": \"abcdefghij\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"message\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/user/password.go",
    "groupTitle": "User",
    "name": "PutApiChange_password",
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/api/user/:id",
    "title": "Change User",
    "description": "<p>Change User</p>",
    "group": "User",
    "version": "0.1.0",
    "permission": [
      {
        "name": "Authorize",
        "title": "Can be access after login",
        "description": ""
      }
    ],
    "examples": [
      {
        "title": "Example usage:",
        "content": "{\n\t\t\"password\": \"ab23fecd32efw34efghe4g5tvij\"\n\t\t\"name\": \"user_test01\",\n\t\t\"email_address\": \"test@mail.com\",\n\t\t\"access_level\": 1\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "id",
            "description": "<p>Page ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Hash Password</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "active",
            "description": "<p>Active Flag</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_logged_in",
            "description": "<p>Last Login Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "last_accessed",
            "description": "<p>Last Access Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "failed_attempts",
            "description": "<p>Failed Login Attempts</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "lock_until",
            "description": "<p>Lock Until Timestamp</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ip_address",
            "description": "<p>IP Address</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "activate_token",
            "description": "<p>Activate Token</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "reminder_token",
            "description": "<p>Reminder Token</p>"
          },
          {
            "group": "Success 200",
            "type": "Int",
            "optional": false,
            "field": "reminder_time",
            "description": "<p>Reminder Timestamp</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Ok\n{\n\t\t\"id\": 1,\n\t\t\"password\": \"ab23fecd32efw34efghe4g5tvij\"\n\t\t\"active\": 1,\n\t\t\"name\": \"user_test01\",\n\t\t\"email_address\": \"test@mail.com\",\n\t\t\"last_logged_in\": null,\n\t\t\"access_level\": 1,\n\t\t\"last_accessed\": null,\n\t\t\"failed_attempts\": null,\n\t\t\"lock_until\": null,\n\t\t\"ip_address\": null,\n\t\t\"activate_token\": null,\n\t\t\"reminder_token\": null,\n\t\t\"reminder_time\": null\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "password",
            "description": "<p>User Password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>User Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email_address",
            "description": "<p>User Email Address</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "access_level",
            "description": "<p>User Access Level</p>"
          }
        ]
      }
    },
    "filename": "./models/user/change.go",
    "groupTitle": "User",
    "name": "PutApiUserId",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Access token.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "{\n\t\t\"Authorization\": \"Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 40X": [
          {
            "group": "Error40X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ],
        "Error 50X": [
          {
            "group": "Error50X",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>Error Message</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 40X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 50X\n{\n\t\t\"message\": \"Error message\"\n}",
          "type": "json"
        }
      ]
    }
  }
] });
