import click
from localseek import vd, vdqdrant
from localseek import chat as _chat

def main():
    cli()

@click.group()
def cli():
    pass

@cli.command()
@click.option('--path', type=str, required=True)
def save(path: str):
    click.echo('[cli] start `save`')
    vd.save(path)

@cli.command()
@click.option('--keyword', type=str, required=True)
def search(keyword: str):
    click.echo('[cli] start `search`')
    vd.search(keyword)

@cli.command()
def chat():
    click.echo('[cli] start `chat`')
    _chat.chat()

# qdrant
@cli.command()
@click.option('--path', type=str, required=True)
def qdrantsave(path: str):
    click.echo('[cli] start `save`')
    vdqdrant.save(path)

@cli.command()
@click.option('--keyword', type=str, required=True)
def qdrantsearch(keyword: str):
    click.echo('[cli] start `search`')
    vdqdrant.search(keyword)
