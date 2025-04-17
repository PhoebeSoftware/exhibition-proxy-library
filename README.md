# exhibtion-proxy

Proxy responsible for fetching IGDB game data.
Why use a proxy? So each user does not have to get their own IGDB client id and secret.

Base URL = localhost:3030

## Usage
Get game data by id
```
GET /game/{igdbid}
```
Example JSON: 
```
{
  "id": 125174,
  "name": "Overwatch 2",
  "summary": "Overwatch 2 is a free-to-play shooter featuring 40+ epic heroes, each with game-changing abilities. Choose your hero, group up with your friends and battle across all-new maps and modes in the ultimate team-based shooter.",
  "cover": {
    "image_id": "co885f"
  },
  "cover_url": "https://images.igdb.com/igdb/image/upload/t_cover_big/co885f.jpg",
  "artworks": [
    {
      "image_id": "ar6cy"
    },
    {
      "image_id": "ar6cz"
    },
    {
      "image_id": "ar6d0"
    },
    {
      "image_id": "ar6d1"
    },
    {
      "image_id": "ar6a9"
    },
    {
      "image_id": "ar6aa"
    },
    {
      "image_id": "ar6ad"
    },
    {
      "image_id": "ar6ac"
    },
    {
      "image_id": "ar6d2"
    },
    {
      "image_id": "ar6ab"
    }
  ],
  "artwork_url_list": [
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6cy.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6cz.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6d0.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6d1.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6a9.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6aa.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6ad.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6ac.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6d2.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/ar6ab.jpg"
  ],
  "screenshots": [
    {
      "image_id": "sc78ey"
    },
    {
      "image_id": "sc78ew"
    },
    {
      "image_id": "sc78ex"
    },
    {
      "image_id": "sc78ev"
    },
    {
      "image_id": "sc78eu"
    },
    {
      "image_id": "sc78ez"
    },
    {
      "image_id": "sc78f0"
    },
    {
      "image_id": "sc78f3"
    },
    {
      "image_id": "sc78f2"
    },
    {
      "image_id": "sc78f1"
    },
    {
      "image_id": "sctkgr"
    },
    {
      "image_id": "sctkgs"
    },
    {
      "image_id": "sctkgt"
    },
    {
      "image_id": "sctkgu"
    },
    {
      "image_id": "sctkgv"
    },
    {
      "image_id": "sctkgw"
    }
  ],
  "screenshot_url_list": [
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78ey.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78ew.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78ex.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78ev.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78eu.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78ez.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78f0.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78f3.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78f2.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sc78f1.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgr.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgs.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgt.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgu.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgv.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sctkgw.jpg"
  ]
}
```

Get game data by name
```
GET /game/?name={name}
```
Example JSON:
```
[
  {
    "id": 261335,
    "name": "Overwatch 2: Hero Collection",
    "summary": "Jumpstart your hero roster and dominate the battlefield",
    "cover": {
      "image_id": "co6y2d"
    },
    "cover_url": "https://images.igdb.com/igdb/image/upload/t_cover_big/co6y2d.jpg",
    "artworks": [
      {
        "image_id": "ar2ggj"
      },
      {
        "image_id": "ar2gh0"
      }
    ],
    "artwork_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/ar2gh0.jpg"
    ],
    "screenshots": [
      {
        "image_id": "scnwbr"
      }
    ],
    "screenshot_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/scnwbr.jpg"
    ]
  },
  {
    "id": 261349,
    "name": "Overwatch 2: Invasion Bundle",
    "summary": "Get started on your mission to save the world with the Overwatch 2: Invasion Bundle!",
    "cover": {
      "image_id": "co84ii"
    },
    "cover_url": "https://images.igdb.com/igdb/image/upload/t_cover_big/co84ii.jpg",
    "artworks": [
      {
        "image_id": "ar2ggy"
      }
    ],
    "artwork_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/ar2ggy.jpg"
    ],
    "screenshots": [
      {
        "image_id": "scq65i"
      }
    ],
    "screenshot_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/scq65i.jpg"
    ]
  },
  {
    "id": 261336,
    "name": "Overwatch 2: Complete Hero Collection",
    "summary": "New to Overwatch 2? Unlock the full game experience with the Overwatch 2: Complete Hero Collection!",
    "cover": {
      "image_id": "co7nne"
    },
    "cover_url": "https://images.igdb.com/igdb/image/upload/t_cover_big/co7nne.jpg",
    "artworks": [
      {
        "image_id": "ar2ggi"
      },
      {
        "image_id": "ar2ggv"
      }
    ],
    "artwork_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/ar2ggv.jpg"
    ],
    "screenshots": [
      {
        "image_id": "scq65n"
      }
    ],
    "screenshot_url_list": [
      "https://images.igdb.com/igdb/image/upload/t_1080p/scq65n.jpg"
    ]
  },
```
