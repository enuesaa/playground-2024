import click
from localseek import vd
from localseek import chat as _chat
from dotenv import load_dotenv
load_dotenv()

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
