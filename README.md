# Fish Fight Back
[![Donate via LiberaPay](https://img.shields.io/liberapay/receives/rocketnine.space.svg?logo=liberapay)](https://liberapay.com/rocketnine.space)
[![Donate via Patreon](https://img.shields.io/badge/dynamic/json?color=%23e85b46&label=Patreon&query=data.attributes.patron_count&suffix=%20patrons&url=https%3A%2F%2Fwww.patreon.com%2Fapi%2Fcampaigns%2F5252223)](https://www.patreon.com/rocketnine)

[Bullet hell](https://en.wikipedia.org/wiki/Shoot_%27em_up#Bullet_hell) video game featuring fishy revenge

This game was created for the [Mini Jam 108](https://itch.io/jam/mini-jam-108-seaside) game jam.

## Play

### Browser

[**Play in your browser**](https://rocketnine.itch.io/fishfightback)

### Compile

**Note:** The following assets are required to compile Fish Fight Back:
- [Cozy Fishing](https://shubibubi.itch.io/cozy-fishing) at `/asset/image/cozy-fishing/`
- [Cozy People](https://shubibubi.itch.io/cozy-people) at `/asset/image/cozy-people/`

These assets are not free, and are not included in this repository.

Install the dependencies listed for [your platform](https://github.com/hajimehoshi/ebiten/blob/main/README.md#platforms),
then run the following command:

`go install code.rocketnine.space/tslocum/fishfightback@latest`

Run `~/go/bin/fishfightback` to play.

## Support

Please share issues and suggestions [here](https://code.rocketnine.space/tslocum/fishfightback/issues).

## Credits

- [Trevor Slocum](https://rocketnine.space) - Game design and programming
- [node punk](https://open.spotify.com/artist/15eFpWQPNRxB89PnFNWvjU?si=z-jfVwYHTxugaC-BGZiyNg) - Music

## Dependencies

- [ebiten](https://github.com/hajimehoshi/ebiten) - Game engine
- [gohan](https://code.rocketnine.space/tslocum/gohan) - Entity Component System framework
