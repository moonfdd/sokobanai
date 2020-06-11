unit sokobanai;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls;

type

  { TFmMain }

  TFmMain = class(TForm)
    BtnClear: TButton;
    BtnSpace: TButton;
    BtnPerson: TButton;
    BtnBox: TButton;
    BtnTarget: TButton;
    BtnGoAI: TButton;
    BtnBackAI: TButton;
    BtnTwoWayAI: TButton;
    BtnPrevious: TButton;
    BtnNext: TButton;
    ImageSelect: TImage;
    Panel1: TPanel;
    PMap: TPanel;
    TxtMsg: TStaticText;
    procedure BtnBackAIClick(Sender: TObject);
    procedure BtnBoxClick(Sender: TObject);
    procedure BtnClearClick(Sender: TObject);
    procedure BtnGoAIClick(Sender: TObject);
    procedure BtnNextClick(Sender: TObject);
    procedure BtnPersonClick(Sender: TObject);
    procedure BtnPreviousClick(Sender: TObject);
    procedure BtnSpaceClick(Sender: TObject);
    procedure BtnTargetClick(Sender: TObject);
    procedure BtnTwoWayAIClick(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure PMapClick(Sender: TObject);
  private

  public

  end;

var
  FmMain: TFmMain;

implementation

{$R *.lfm}

{ TFmMain }

procedure TFmMain.FormCreate(Sender: TObject);
begin
  //
end;

procedure TFmMain.BtnSpaceClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnTargetClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnTwoWayAIClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnPersonClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnPreviousClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnBoxClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnClearClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnBackAIClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnGoAIClick(Sender: TObject);
begin

end;

procedure TFmMain.BtnNextClick(Sender: TObject);
begin

end;

procedure TFmMain.PMapClick(Sender: TObject);
begin

end;

end.

