# Go + HTMX + Bulma CSS

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

A little project to learn templating with Go using Bulma CSS

## API Reference

#### Get all items

```http
  POST /add-films
```

| Parameter  | Type     | Description                        |
| :--------- | :------- | :--------------------------------- |
| `title`    | `string` | **Required**. title of the film    |
| `director` | `string` | **Required**. director of the film |

