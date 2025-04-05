# exhibtion-proxy

Proxy responsible for fetching IGDB game data.
Why use a proxy? So each user does not have to get their own IGDB client id and secret.

Base URL = localhost:3030

## Usage
Get game data by id
```
GET /getGame/{igdbid}
```
Example JSON: 
```
{
  "id": 11544,
  "name": "Paladins",
  "summary": "Join 25+ million players in Paladins, the free-to-play fantasy team-based shooter sensation. Wield guns and magic as a legendary Champion of the Realm, customizing your core set of abilities to play exactly how you want to play.\n\nPaladins is set in a vibrant fantasy world and features a diverse cast of Champions ranging from sharpshooting humans to mech-riding goblins, mystical elves, and jetpack-clad dragons. Each Champion brings a unique set of abilities to the battlefield and new Champions are regularly added to Paladins, keeping the game exciting.",
  "cover": {
    "image_id": "co1p3u"
  },
  "cover_url": "https://images.igdb.com/igdb/image/upload/t_cover_big/co1p3u.jpg",
  "artworks": [
    {
      "image_id": "arbi7"
    }
  ],
  "artwork_url_list": [
    "https://images.igdb.com/igdb/image/upload/t_1080p/arbi7.jpg"
  ],
  "screenshots": [
    {
      "image_id": "bl5wghaei66cmhcdwpd7"
    },
    {
      "image_id": "sco4sf"
    },
    {
      "image_id": "sco4sd"
    },
    {
      "image_id": "sco4sg"
    },
    {
      "image_id": "sco4sh"
    },
    {
      "image_id": "sco4se"
    },
    {
      "image_id": "sco4si"
    },
    {
      "image_id": "sco4sj"
    },
    {
      "image_id": "sco4sk"
    }
  ],
  "screenshot_url_list": [
    "https://images.igdb.com/igdb/image/upload/t_1080p/bl5wghaei66cmhcdwpd7.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sf.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sd.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sg.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sh.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4se.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4si.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sj.jpg",
    "https://images.igdb.com/igdb/image/upload/t_1080p/sco4sk.jpg"
  ]
}
```

Get game data by name
```
SOON TM
```
